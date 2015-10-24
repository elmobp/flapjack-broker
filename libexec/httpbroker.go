package main

import (
	"encoding/json"
	"flapjackconfig"
	"flapjackbroker"
	"flapjackfeeder"
	"fmt"
	"github.com/go-martini/martini"
	"gopkg.in/alecthomas/kingpin.v1"
	"log"
	"net/http"
	"os"
	"time"
)

// cacheState stores a cache of event state to be sent to Flapjack.
// The event state is queried later when submitting events periodically
// to Flapjack.
func cacheState(updates chan flapjackconfig.State, state map[string]flapjackconfig.State) {
	for ns := range updates {
		key := ns.Entity + ":" + ns.Check
		state[key] = ns
	}
}

// submitCachedState periodically samples the cached state, sends it to Flapjack.
func submitCachedState(states map[string]flapjackconfig.State, config Config) {
	transport, err := flapjackbroker.Dial(config.Server, config.Database)
	broker := []interface{}{}
	broker = append(broker, "httpbroker")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	for {
		log.Printf("Number of cached states: %d\n", len(states))
		for id, state := range states {
			now := time.Now().Unix()
			event := flapjackbroker.Event{
				Entity:  state.Entity,
				Check:   state.Check,
				Type:    state.Type,
				State:   state.State,
				Summary: state.Summary,
				Time:    now,
				Tags:    broker,
			}

			if config.Debug {
				log.Printf("Sending event data for %s\n", id)
			}
			transport.Send(event)
		}
		time.Sleep(config.Interval)
	}
}

var (
	port     = kingpin.Flag("port", "Address to bind HTTP server (default 3090)").Default("3090").OverrideDefaultFromEnvar("PORT").String()
	server   = kingpin.Flag("server", "Redis server to connect to (default localhost:6380)").Default("localhost:6380").String()
	database = kingpin.Flag("database", "Redis database to connect to (default 0)").Int() // .Default("13").Int()
	interval = kingpin.Flag("interval", "How often to submit events (default 10s)").Default("10s").Duration()
	debug    = kingpin.Flag("debug", "Enable verbose output (default false)").Bool()
)

type Config struct {
	Port     string
	Server   string
	Database int
	Interval time.Duration
	Debug    bool
}

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	config := Config{
		Server:   *server,
		Database: *database,
		Interval: *interval,
		Debug:    *debug,
		Port:     ":" + *port,
	}
	if config.Debug {
		log.Printf("Booting with config: %+v\n", config)
	}

	updates := make(chan flapjackconfig.State)
	state := map[string]flapjackconfig.State{}

	go cacheState(updates, state)
	go submitCachedState(state, config)

	m := martini.Classic()
	// Handle SNS
	m.Group("/state", func(r martini.Router) {
		r.Post("", func(res http.ResponseWriter, req *http.Request) {
			flapjackfeeder.CreateCloudwatchState(updates, res, req)
		})
		r.Get("", func() []byte {
			data, _ := json.Marshal(state)
			return data
		})
	})
	// Handle Flapjack
	m.Group("/flapjack", func(r martini.Router) {
		r.Post("", func(res http.ResponseWriter, req *http.Request) {
			flapjackfeeder.CreateFlapjackState(updates, res, req)
		})
		r.Get("", func() []byte {
			data, _ := json.Marshal(state)
			return data
		})
	})
	// Handle New Relic
	m.Group("/newrelic", func(r martini.Router) {
		r.Post("", func(res http.ResponseWriter, req *http.Request) {
			flapjackfeeder.CreateNewRelicState(updates, res, req)
		})
		r.Get("", func() []byte {
			data, _ := json.Marshal(state)
			return data
		})
	})
	log.Fatal(http.ListenAndServe(config.Port, m))
}

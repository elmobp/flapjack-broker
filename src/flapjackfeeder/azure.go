package flapjackfeeder

import (
        "encoding/json"
        "flapjackbroker"
        "flapjackconfig"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
	"strings"
        "time"
)

func CreateAzureState(updates chan flapjackconfig.State, w http.ResponseWriter, r *http.Request) {
	var AzureAlert flapjackconfig.AzureAlarm
	var event_state string
	var state flapjackconfig.State

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		message := "Error: Couldn't read request body: %s\n"
		log.Printf(message, err)
		fmt.Fprintf(w, message, err)
		return
	}

	// Check if its a New Relic event
	err = json.Unmarshal(body, &AzureAlert)
	if err != nil {
		message := "Error: Couldn't Decode New Relic into JSON: %s\n"
		log.Println(message, err)
		fmt.Fprintf(w, message, err)
		return
	}

	switch strings.ToLower(AzureAlert.Context.Condition.Operator) {
	case "GreaterThan":
		event_state = "critical"
	default:
		event_state = "ok"
	}
	new_details := fmt.Sprint("Azure Alert Received: ", AzureAlert.Context.Condition.MetricName, " Current value: ", AzureAlert.Context.Condition.MetricValue)
	tmpclient := fmt.Sprint(AzureAlert.Context.Name, "-bpdyn")
	state = flapjackconfig.State{
		flapjackbroker.Event{
			Entity: tmpclient,
			Check:  AzureAlert.Context.Description,
			// Type:    "service", // @TODO: Make this magic
			State:   event_state,
			Summary: new_details,
		},
		0,
	}

	// Populate a time if none has been set.
	if state.Time == 0 {
		state.Time = time.Now().Unix()
	}

	if len(state.Type) == 0 {
		state.Type = "service"
	}

	if state.TTL == 0 {
		state.TTL = 300
	}

	updates <- state

	json, _ := json.Marshal(state)
	message := "Caching state: %s\n"
	log.Printf(message, json)
	fmt.Fprintf(w, message, json)
}




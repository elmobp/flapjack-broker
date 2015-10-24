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

func CreateNewRelicState(updates chan flapjackconfig.State, w http.ResponseWriter, r *http.Request) {
	var NRAlarm flapjackconfig.NewRelicAlert
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
	err = json.Unmarshal(body, &NRAlarm)
	if err != nil {
		message := "Error: Couldn't Decode New Relic into JSON: %s\n"
		log.Println(message, err)
		fmt.Fprintf(w, message, err)
		return
	}

	switch strings.ToLower(NRAlarm.CurrentState) {
	case "alarm":
		event_state = "critical"
	default:
		event_state = "ok"
	}
	new_details := fmt.Sprint("New Relic Alert Received: ", NRAlarm.Details, " Ack the alert using the following URL: ", NRAlarm.IncidentAcknowledgeURL)
	state = flapjackconfig.State{
		flapjackbroker.Event{
			Entity: NRAlarm.PolicyName,
			Check:  NRAlarm.ConditionName,
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




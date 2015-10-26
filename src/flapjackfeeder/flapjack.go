package flapjackfeeder

import (
        "encoding/json"
        "flapjackbroker"
        "flapjackconfig"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "time"
)

func CreateFlapjackState(updates chan flapjackconfig.State, w http.ResponseWriter, r *http.Request) {
	var input_state flapjackconfig.InputState
	var state flapjackconfig.State
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		message := "Error: Couldn't read request body: %s\n"
		log.Printf(message, err)
		fmt.Fprintf(w, message, err)
		return
	}

	// Check if its a Flapjack
	err = json.Unmarshal(body, &input_state)
	if err != nil {
		message := "Error: Couldn't Decode Flapjack into JSON: %s\n"
		log.Println(message, err)
		fmt.Fprintf(w, message, err)
		return
	}
        tmpclient := fmt.Sprint(input_state.Entity, "-bpdyn")
	state = flapjackconfig.State{
		flapjackbroker.Event{
			Entity: tmpclient,
			Check:  input_state.Check,
			// Type:    "service", // @TODO: Make this magic
			State:   input_state.State,
			Summary: input_state.Summary,
			Time:    input_state.Time,
		},
		input_state.TTL,
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



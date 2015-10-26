package  flapjackfeeder

import (
	"encoding/json"
	"flapjackbroker"
	"flapjackconfig"
	"strings"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// handler caches
func CreateCloudwatchState(updates chan flapjackconfig.State, w http.ResponseWriter, r *http.Request) {
	var SNSData flapjackconfig.SNSNotification
	var sns_subscription flapjackconfig.SNSSubscribe
	var cw_alarm flapjackconfig.CWAlarm
	var state flapjackconfig.State
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		message := "Error: Couldn't read request body: %s\n"
		log.Printf(message, err)
		fmt.Fprintf(w, message, err)
		return
	}

	// Check if its a SNS
	err = json.Unmarshal(body, &SNSData)
	if err != nil {
		message := "Error: Couldn't Decode SNS into JSON: %s\n"
		log.Println(message, err)
		fmt.Fprintf(w, message, err)
		return
	}

	json.Unmarshal(body, &sns_subscription)
	if sns_subscription.SubscribeURL != "" {
		http.Get(sns_subscription.SubscribeURL)
		return
	}

	input_message := []byte(SNSData.Message)
	err = json.Unmarshal(input_message, &cw_alarm)
	if err != nil {
		message := "Error: Couldn't read request body from the SNS message: %s\n"
		log.Println(message, err)
		fmt.Fprintf(w, message, err)
		return
	}

	var event_state string
	switch strings.ToLower(cw_alarm.NewStateValue) {
	case "alarm":
		event_state = "critical"
	default:
		event_state = "ok"
	}
	s := strings.Split(cw_alarm.AlarmDescription, ":")
        tmpclient := fmt.Sprint(s[0], "-bpdyn")
	tmpdata := fmt.Sprint(s[1], " ", cw_alarm.Trigger.MetricName)
	state = flapjackconfig.State{
		flapjackbroker.Event{
			Entity:  tmpclient,
			Check:   tmpdata,
			State:   event_state,
			Summary: cw_alarm.NewStateReason,
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


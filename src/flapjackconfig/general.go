package flapjackconfig

import (
 "flapjackbroker"
)
type EventFormat int

const (
	Normal   EventFormat = 1
	SNS      EventFormat = 2
	Newrelic EventFormat = 3
)

// State is a basic representation of a Flapjack event, with some extra field.
// The extra fields handle state expiry.
// Find more at http://flapjack.io/docs/1.0/development/DATA_STRUCTURES
type State struct {
	flapjackbroker.Event
	TTL int64 `json:"ttl"`
}

type InputState struct {
	flapjackbroker.Event
	TTL int64 `json:"ttl"`
}

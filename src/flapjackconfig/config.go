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
type SNSSubscribe struct {
	Message          string `json:"Message"`
	MessageID        string `json:"MessageId"`
	Signature        string `json:"Signature"`
	SignatureVersion string `json:"SignatureVersion"`
	SigningCertURL   string `json:"SigningCertURL"`
	SubscribeURL     string `json:"SubscribeURL"`
	Timestamp        string `json:"Timestamp"`
	Token            string `json:"Token"`
	TopicArn         string `json:"TopicArn"`
	Type             string `json:"Type"`
}
type SNSNotification struct {
	Message          string `json:"Message"`
	MessageID        string `json:"MessageId"`
	Signature        string `json:"Signature"`
	SignatureVersion string `json:"SignatureVersion"`
	SigningCertURL   string `json:"SigningCertURL"`
	Subject          string `json:"Subject"`
	Timestamp        string `json:"Timestamp"`
	TopicArn         string `json:"TopicArn"`
	Type             string `json:"Type"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}
type CWAlarm struct {
	AWSAccountID     string      `json:"AWSAccountId"`
	AlarmDescription interface{} `json:"AlarmDescription"`
	AlarmName        string      `json:"AlarmName"`
	NewStateReason   string      `json:"NewStateReason"`
	NewStateValue    string      `json:"NewStateValue"`
	OldStateValue    string      `json:"OldStateValue"`
	Region           string      `json:"Region"`
	StateChangeTime  string      `json:"StateChangeTime"`
	Time             int64       `json:"Time"`
	Trigger          struct {
		ComparisonOperator string `json:"ComparisonOperator"`
		Dimensions         []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"Dimensions"`
		EvaluationPeriods int         `json:"EvaluationPeriods"`
		MetricName        string      `json:"MetricName"`
		Namespace         string      `json:"Namespace"`
		Period            int         `json:"Period"`
		Statistic         string      `json:"Statistic"`
		Threshold         float64     `json:"Threshold"`
		Unit              interface{} `json:"Unit"`
	} `json:"Trigger"`
}

type NewRelicAlert struct {
	AccountID              int    `json:"account_id"`
	AccountName            string `json:"account_name"`
	ConditionID            int    `json:"condition_id"`
	ConditionName          string `json:"condition_name"`
	CurrentState           string `json:"current_state"`
	Details                string `json:"details"`
	EventType              string `json:"event_type"`
	IncidentAcknowledgeURL string `json:"incident_acknowledge_url"`
	IncidentID             int    `json:"incident_id"`
	IncidentURL            string `json:"incident_url"`
	Owner                  string `json:"owner"`
	PolicyName             string `json:"policy_name"`
	PolicyURL              string `json:"policy_url"`
	RunbookURL             string `json:"runbook_url"`
	Severity               string `json:"severity"`
	Targets                []struct {
		ID     string `json:"id"`
		Labels struct {
			Label string `json:"label"`
		} `json:"labels"`
		Link    string `json:"link"`
		Name    string `json:"name"`
		Product string `json:"product"`
		Type    string `json:"type"`
	} `json:"targets"`
	Timestamp int `json:"timestamp"`
}

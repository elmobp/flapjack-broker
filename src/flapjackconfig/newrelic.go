package flapjackconfig

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


package models

// NewRelics is a list of NewRelic integrations
type NewRelics struct {
	Data       []NewRelic `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// NewRelic is a NewRelic integration
type NewRelic struct {
	Data
	Attributes NewRelicAttributes `json:"attributes"`
}

type NewRelicAttributes struct {
	Call           bool   `json:"call"`
	Sms            bool   `json:"sms"`
	Email          bool   `json:"email"`
	TeamWait       int    `json:"team_wait"`
	RecoveryPeriod any    `json:"recovery_period"`
	PolicyID       any    `json:"policy_id"`
	Push           bool   `json:"push"`
	Name           string `json:"name"`
	Paused         bool   `json:"paused"`
	AlertingRule   string `json:"alerting_rule"`
	WebhookURL     string `json:"webhook_url"`
}

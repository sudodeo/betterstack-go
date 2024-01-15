package models

// Heartbeats represents a response containing a list of Heartbeats
type Heartbeats struct {
	Data       []Heartbeat
	Pagination pagination
}

// Heartbeat represents a Heartbeat
type Heartbeat struct {
	ID         string
	Type       string
	Attributes heartbeatAttributes
}

type heartbeatAttributes struct {
	URL              string `json:"url"`
	Name             string `json:"name"`               // The name of the service for this heartbeat.
	Period           int    `json:"period"`             // How often should we expect this heartbeat? In seconds Minimum value: 30 seconds
	Grace            int    `json:"grace"`              // Heartbeats can fluctuate; specify this value to control what is still acceptable Minimum value: 0 seconds We recommend setting this to approx. 20% of period
	Call             bool   `json:"call"`               // Should we call the on-call person?
	Sms              bool   `json:"sms"`                // Should we send an SMS to the on-call person?
	Email            bool   `json:"email"`              // Should we send an email to the on-call person?
	Push             bool   `json:"push"`               // Should we send a push notification to the on-call person?
	TeamWait         int    `json:"team_wait"`          // How long should we wait before escalating the incident alert to the team? Leave blank to disable escalating to the entire team.
	HeartbeatGroupID int    `json:"heartbeat_group_id"` // Set this attribute if you want to add this heartbeat to a heartbeat group
	SortIndex        int    `json:"sort_index"`         // An index controlling the position of a heartbeat in the heartbeat group.
	PausedAt         string `json:"paused_at"`          // Set to true to pause monitoring — we won't notify you about downtime. Set to false to resume monitoring
	CreatedAt        string `json:"created_at"`         //
	UpdatedAt        string `json:"updated_at"`
	Status           string `json:"status"`
}

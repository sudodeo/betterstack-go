package models

// HeartbeatReqBody represents the request body for heartbeats
type HeartbeatReqBody struct {
	Name             string `json:"name"`                         // Required. The name of the service for this heartbeat.
	Period           int    `json:"period"`                       // Required. How often should we expect this heartbeat? In seconds Minimum value: 30 seconds
	Grace            int    `json:"grace"`                        // Required. Heartbeats can fluctuate; specify this value to control what is still acceptable Minimum value: 0 seconds We recommend setting this to approx. 20% of period
	Call             bool   `json:"call,omitempty"`               // Should we call the on-call person?
	Sms              bool   `json:"sms,omitempty"`                // Should we send an SMS to the on-call person?
	Email            bool   `json:"email,omitempty"`              // Should we send an email to the on-call person?
	Push             bool   `json:"push,omitempty"`               // Should we send a push notification to the on-call person?
	TeamWait         int    `json:"team_wait,omitempty"`          // How long should we wait before escalating the incident alert to the team? Leave blank to disable escalating to the entire team.
	HeartbeatGroupID int    `json:"heartbeat_group_id,omitempty"` // Set this attribute if you want to add this heartbeat to a heartbeat group
	SortIndex        int    `json:"sort_index,omitempty"`         // An index controlling the position of a heartbeat in the heartbeat group.
	PausedAt         string `json:"paused_at,omitempty"`          // Set to true to pause monitoring — we won't notify you about downtime. Set to false to resume monitoring
	PolicyID         string `json:"policy_id,omitempty"`          // Set the escalation policy for the monitor.
}

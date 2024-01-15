package models

// IncidentReqBody represents the request body for an incident request.
type IncidentReqBody struct {
	// Required: RequesterEmail is the e-mail of the user who requested the incident.
	RequesterEmail string `json:"requester_email,omitempty"`

	// Required: Summary is a brief summary of the incident.
	Summary string `json:"summary"`

	// Name is the short name of the incident.
	Name string `json:"name,omitempty"`

	// Description is the full description of the incident.
	Description string `json:"description"`

	// Call indicates whether we should call the on-call person.
	Call bool `json:"call,omitempty"`

	// SMS indicates whether we should send an SMS to the on-call person.
	SMS bool `json:"sms,omitempty"`

	// Email indicates whether we should send an email to the on-call person.
	Email bool `json:"email,omitempty"`

	// Push indicates whether we should send a push notification to the on-call person.
	Push bool `json:"push,omitempty"`

	// TeamWait is how long to wait before escalating the incident alert to the team.
	// Leave blank to disable escalating to the entire team. In seconds.
	TeamWait int `json:"team_wait,omitempty"`

	// PolicyID is the ID of the escalation policy with which you'd like to escalate this incident.
	PolicyID string `json:"policy_id,omitempty"`
}

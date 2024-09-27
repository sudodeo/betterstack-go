package models

type PolicyReqBody struct {
	TeamName    string       `json:"team_name,omitempty"`
	Name        string       `json:"name"`
	RepeatCount int          `json:"repeat_count,omitempty"`
	RepeatDelay int          `json:"repeat_delay,omitempty"`
	Steps       []PolicyStep `json:"steps,omitempty"`
}

type EscalationPolicyAttributes struct {
	Name          string       `json:"name,omitempty"`
	RepeatCount   int          `json:"repeat_count,omitempty"`
	RepeatDelay   int          `json:"repeat_delay,omitempty"`
	IncidentToken string       `json:"incident_token,omitempty"`
	PolicyGroupID int          `json:"policy_group_id,omitempty"`
	TeamName      string       `json:"team_name,omitempty"`
	Steps         []PolicyStep `json:"steps,omitempty"`
}

type PolicyStep struct {
	Type           string             `json:"type,omitempty"`
	WaitBefore     int                `json:"wait_before,omitempty"`
	UrgencyID      int                `json:"urgency_id,omitempty"`
	StepMembers    []PolicyStepMember `json:"step_members,omitempty"`
	Timezone       string             `json:"timezone,omitempty"`
	Days           []string           `json:"days,omitempty"`
	TimeFrom       string             `json:"time_from,omitempty"`
	TimeTo         string             `json:"time_to,omitempty"`
	MetadataKey    string             `json:"metadata_key,omitempty"`
	MetadataValues []string           `json:"metadata_values,omitempty"`
	PolicyID       int                `json:"policy_id"`
}

type PolicyStepMember struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Email string `json:"email,omitempty"`
}

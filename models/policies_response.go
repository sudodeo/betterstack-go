package models

type EscalationPolicies struct {
	Data       []EscalationPolicy `json:"data"`
	Pagination pagination         `json:"pagination"`
}

type EscalationPolicy struct {
	ID         string                     `json:"id"`
	Type       string                     `json:"type"`
	Attributes EscalationPolicyAttributes `json:"attributes,omitempty"`
}

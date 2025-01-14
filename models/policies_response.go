package models

type EscalationPolicies struct {
	Data       []EscalationPolicy `json:"data"`
	Pagination Pagination         `json:"pagination"`
}

type EscalationPolicy struct {
	ID         string                     `json:"id"`
	Type       string                     `json:"type"`
	Attributes EscalationPolicyAttributes `json:"attributes,omitempty"`
}

package models

// Groups represents the response body for monitor groups and heartbeat groups
type Groups struct {
	Data       []Group    `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// Group represents the a monitor group or heartbeat group
type Group struct {
	ID         string          `json:"id,omitempty"`
	Type       string          `json:"type,omitempty"`
	Attributes GroupAttributes `json:"attributes,omitempty"`
}

type GroupAttributes struct {
	Name      string `json:"name,omitempty"`
	SortIndex int    `json:"sort_index,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Paused    bool   `json:"paused,omitempty"`
}

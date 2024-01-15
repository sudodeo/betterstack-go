package models

// GroupReqBody represents the request body for monitor groups and heartbeat groups
type GroupReqBody struct {
	Name      string `json:"name"`
	Paused    bool   `json:"paused,omitempty"`     // optional,defaults to false
	SortIndex *int   `json:"sort_index,omitempty"` // optional, can be null
}

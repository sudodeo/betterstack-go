package models

// StatusUpdates represents a collection of status updates with pagination information.
type StatusUpdates struct {
	Data       []StatusUpdate `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

// StatusUpdate represents an individual status update with its attributes.
type StatusUpdate struct {
	Data
	Attributes IncludedAttributes `json:"attributes"`
}

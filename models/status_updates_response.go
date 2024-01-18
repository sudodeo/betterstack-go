package models

// StatusUpdates represents a collection of status updates with pagination information.
type StatusUpdates struct {
	Data       []StatusUpdate
	Pagination pagination
}

// StatusUpdate represents an individual status update with its attributes.
type StatusUpdate struct {
	data
	Attributes includedAttributes
}

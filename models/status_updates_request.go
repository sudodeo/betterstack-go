package models

// StatusUpdateReqBody represents the request body for a status report update.
type StatusUpdateReqBody struct {
	AffectedResources []AffectedResource `json:"affected_resources,omitempty"` // Required
	Message           string             `json:"message,omitempty"`            // Required
	PublishedAt       string             `json:"published_at,omitempty"`
}

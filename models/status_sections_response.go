package models

// StatusPageSections represents a collection of status pages.
type StatusPageSections struct {
	Data       []StatusPageSection `json:"data,omitempty"`       // Data contains the list of status pages.
	Pagination Pagination          `json:"pagination,omitempty"` // Pagination contains the pagination information.
}

// StatusPageSection represents a status page object.
type StatusPageSection struct {
	ID         string                      `json:"id,omitempty"`
	Type       string                      `json:"type,omitempty"`
	Attributes StatusPageSectionAttributes `json:"attributes,omitempty"`
}

type StatusPageSectionAttributes struct {
	StatusPageID int    `json:"status_page_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Position     int    `json:"position,omitempty"`
}

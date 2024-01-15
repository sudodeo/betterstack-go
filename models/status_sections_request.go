package models

// StatusPageSectionReqBody represents the request body for creating/updating a status page.
type StatusPageSectionReqBody struct {
	Name     string `json:"name,omitempty"`     // leave blank to hide section header
	Position int    `json:"position,omitempty"` // The position of this resource on your status page, indexed from zero. If you don't specify a position, we add the resource to the end of the status page. When you specify a position of an existing resource, we add the resource to this position and shift resources below to accommodate.
}

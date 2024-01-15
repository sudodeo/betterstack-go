package models

// StatusPageResourceReqBody represents the request body for adding a resource to a status page.
type StatusPageResourceReqBody struct {
	// Required: The ID of the resource you are adding
	ResourceID string `json:"resource_id"`

	// Required: The type of the resource you are adding. Available values: Monitor, Heartbeat, WebhookIntegration, EmailIntegration
	ResourceType string `json:"resource_type"`

	// Required: The resource name displayed publicly on your status page.
	PublicName string `json:"public_name"`

	// The ID of the section which should contain this resource. When omitted, defaults to the first section on the status page.
	StatusPageSectionID string `json:"status_page_section_id,omitempty"`

	// What widget to display for this resource. Expects one of three values: plain - only display status history - display detailed historical status response_times - add a response times chart (only for Monitor . This takes preference over history when both parameters are present.
	WidgetType string `json:"widget_type,omitempty"`

	// The resource name displayed publicly on your status page.
	Explanation string `json:"explanation,omitempty"`

	// The position of this resource on your status page, indexed from zero. If you don't specify a position, we add the resource to the end of the status page. When you specify a position of an existing resource, we add the resource to this position and shift resources below to accommodate.
	Position int `json:"position,omitempty"`
}

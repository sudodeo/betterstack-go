package uptime

// StatusPageResources represents the response containing status page resources
type StatusPageResources struct {
	Data       []StatusPageResource
	Pagination pagination
}

// StatusPageResource represents a status page resource.
type StatusPageResource struct {
	data
	Attributes resourceAttributes `json:"attributes,omitempty"`
}

type resourceAttributes struct {
	StatusPageSectionID int             `json:"status_page_section_id"`
	ResourceID          int             `json:"resource_id"`
	ResourceType        string          `json:"resource_type"`
	History             bool            `json:"history"`
	WidgetType          string          `json:"widget_type"`
	PublicName          string          `json:"public_name"`
	Explanation         string          `json:"explanation"`
	Position            int             `json:"position"`
	Availability        float64         `json:"availability"`
	StatusHistory       []statusHistory `json:"status_history"`
}

type statusHistory struct {
	Day                 string `json:"day,omitempty"`
	Status              string `json:"status,omitempty"`
	DowntimeDuration    int    `json:"downtime_duration,omitempty"`
	MaintenanceDuration string
}

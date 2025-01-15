package models

// StatusPageReports represents a response body for a list of status page reports
type StatusPageReports struct {
	Data       []StatusPageReport `json:"data,omitempty"`
	Pagination Pagination         `json:"pagination,omitempty"`
}

type StatusPageReportData struct {
	Data
	Attributes    reportAttributes `json:"attributes,omitempty"`
	Relationships Relationships    `json:"relationships,omitempty"`
}

// StatusPageReport represents a status report for a status page.
type StatusPageReport struct {
	Data     StatusPageReportData
	Included []Included `json:"included,omitempty"`
}

type reportAttributes struct {
	Title             string             `json:"title,omitempty"`
	ReportType        string             `json:"report_type,omitempty"`
	StartsAt          string             `json:"starts_at,omitempty"`
	EndsAt            string             `json:"ends_at,omitempty"`
	StatusPageID      int                `json:"status_page_id,omitempty"`
	AffectedResources []AffectedResource `json:"affected_resources,omitempty"`
	AggregateState    string             `json:"aggregate_state,omitempty"`
}

type Included struct {
	Data
	Attributes IncludedAttributes `json:"attributes"`
}

type IncludedAttributes struct {
	Message           string             `json:"message"`
	PublishedAt       string             `json:"published_at"`
	StatusReportID    int                `json:"status_report_id"`
	AffectedResources []AffectedResource `json:"affected_resources"`
}

// AffectedResource represents a resource that is affected by a status report.
type AffectedResource struct {
	StatusPageResourceID string `json:"status_page_resource_id,omitempty"`
	Status               string `json:"status,omitempty"` // [resolved, degraded, downtime, maintenance(if resource_type is maintenance)]
}

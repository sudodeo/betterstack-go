package uptime

// StatusPageReportReqBody represents the request body for creating a status page report.
type StatusPageReportReqBody struct {
	Title             string             `json:"title"`                        // Required
	AffectedResources []AffectedResource `json:"affected_resources,omitempty"` // Required to create report
	Message           string             `json:"message,omitempty"`            // Required to create report
	ReportType        string             `json:"report_type,omitempty"`        // manual or maintenance. Defaults to manual
	PublishedAt       string             `json:"published_at,omitempty"`
	StartsAt          string             `json:"starts_at,omitempty"`
	EndsAt            string             `json:"ends_at,omitempty"`
}

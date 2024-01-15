package models

// Incidents represents a collection of incidents.
type Incidents struct {
	Data       []Incident `json:"data,omitempty"`
	Pagination pagination `json:"pagination,omitempty"`
}

// Incident represents an incident.
type Incident struct {
	ID            string             `json:"id,omitempty"`
	Type          string             `json:"type,omitempty"`
	Attributes    incidentAttributes `json:"attributes,omitempty"`
	Relationships relationships      `json:"relationships,omitempty"`
}

type incidentAttributes struct {
	Name               string   `json:"name,omitempty"`
	URL                string   `json:"url,omitempty"`
	HTTPMethod         string   `json:"http_method,omitempty"`
	Cause              string   `json:"cause,omitempty"`
	IncidentGroupID    int      `json:"incident_group_id,omitempty"`
	StartedAt          string   `json:"started_at,omitempty"`
	AcknowledgedAt     string   `json:"acknowledged_at,omitempty"`
	AcknowledgedBy     string   `json:"acknowledged_by,omitempty"`
	ResolvedAt         string   `json:"resolved_at,omitempty"`
	ResolvedBy         string   `json:"resolved_by,omitempty"`
	ResponseContent    string   `json:"response_content,omitempty"`
	ResponseOptions    string   `json:"response_options,omitempty"`
	Regions            []string `json:"regions,omitempty"`
	ResponseURL        string   `json:"response_url,omitempty"`
	ScreenshotURL      string   `json:"screenshot_url,omitempty"`
	OriginURL          string   `json:"origin_url,omitempty"`
	EscalationPolicyID string   `json:"escalation_policy_id,omitempty"`
	Call               bool     `json:"call,omitempty"`
	SMS                bool     `json:"sms,omitempty"`
	Email              bool     `json:"email,omitempty"`
	Push               bool     `json:"push,omitempty"`
}

// IncidentTimelineEvents represents a list of timeline events associated with an incident.
type IncidentTimelineEvents struct {
	Data []timelineItem `json:"data,omitempty"`
}

type timelineItem struct {
	ID         string                 `json:"id,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Attributes timelineItemAttributes `json:"attributes,omitempty"`
}

type timelineItemAttributes struct {
	ItemType string       `json:"item_type,omitempty"`
	At       string       `json:"at,omitempty"`
	Data     timelineData `json:"data,omitempty"`
}

type timelineData struct {
	Title       string       `json:"title,omitempty"`
	Content     string       `json:"content,omitempty"`
	Attachments []attachment `json:"attachments,omitempty"`
	Links       []string     `json:"links,omitempty"`
}

type content struct {
	Text string `json:"text,omitempty"`
	User user   `json:"user,omitempty"`
}

type attachment struct {
	Size     int    `json:"size,omitempty"`
	Filename string `json:"filename,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	URL      string `json:"url,omitempty"`
}

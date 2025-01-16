package models

// Incidents represents a collection of incidents.
type Incidents struct {
	Data       []Incident `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// Incident represents an incident.
type Incident struct {
	ID            string             `json:"id,omitempty"`
	Type          string             `json:"type,omitempty"`
	Attributes    IncidentAttributes `json:"attributes,omitempty"`
	Relationships Relationships      `json:"relationships,omitempty"`
}

type IncidentAttributes struct {
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
	Data []TimelineItem `json:"data,omitempty"`
}

type TimelineItem struct {
	ID         string                 `json:"id,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Attributes TimelineItemAttributes `json:"attributes,omitempty"`
}

type TimelineItemAttributes struct {
	ItemType string       `json:"item_type,omitempty"`
	At       string       `json:"at,omitempty"`
	Data     TimelineData `json:"data,omitempty"`
}

type TimelineData struct {
	Title       string       `json:"title,omitempty"`
	Content     string       `json:"content,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Links       []string     `json:"links,omitempty"`
}

type Content struct {
	Text string `json:"text,omitempty"`
	User User   `json:"user,omitempty"`
}

type Attachment struct {
	Size     int    `json:"size,omitempty"`
	Filename string `json:"filename,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	URL      string `json:"url,omitempty"`
}

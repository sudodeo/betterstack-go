package uptime

// StatusPages represents a collection of status pages.
type StatusPages struct {
	Data       []StatusPage `json:"data,omitempty"`       // Data contains the list of status pages.
	Pagination pagination   `json:"pagination,omitempty"` // Pagination contains the pagination information.
}

// StatusPage represents a status page object.
type StatusPage struct {
	ID         string               `json:"id,omitempty"`
	Type       string               `json:"type,omitempty"`
	Attributes statusPageAttributes `json:"attributes,omitempty"`
}

type statusPageAttributes struct {
	CompanyName              string `json:"company_name,omitempty"`
	CompanyURL               string `json:"company_url,omitempty"`
	ContactURL               string `json:"contact_url,omitempty"`
	LogoURL                  string `json:"logo_url,omitempty"`
	Timezone                 string `json:"timezone,omitempty"`
	Subdomain                string `json:"subdomain,omitempty"`
	CustomDomain             string `json:"custom_domain,omitempty"`
	CustomCSS                string `json:"custom_css,omitempty"`
	GoogleAnalyticsID        string `json:"google_analytics_id,omitempty"`
	MinIncidentLength        int    `json:"min_incident_length,omitempty"`
	Announcement             string `json:"announcement,omitempty"`
	AnnouncementEmbedVisible bool   `json:"announcement_embed_visible,omitempty"`
	AnnouncementEmbedCSS     string `json:"announcement_embed_css,omitempty"`
	AnnouncementEmbedLink    string `json:"announcement_embed_link,omitempty"`
	AutomaticReports         bool   `json:"automatic_reports,omitempty"`
	Subscribable             bool   `json:"subscribable,omitempty"`
	HideFromSearchEngines    bool   `json:"hide_from_search_engines,omitempty"`
	PasswordEnabled          bool   `json:"password_enabled,omitempty"`
	History                  int    `json:"history,omitempty"`
	AggregateState           string `json:"aggregate_state,omitempty"`
	Design                   string `json:"design,omitempty"`
	Theme                    string `json:"theme,omitempty"`
	Layout                   string `json:"layout,omitempty"`
	CreatedAt                string `json:"created_at,omitempty"`
	UpdatedAt                string `json:"updated_at,omitempty"`
}

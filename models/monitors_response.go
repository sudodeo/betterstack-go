package models

// Monitors represents a response containing a list of responses
type Monitors struct {
	Data       []Monitor  `json:"data"`
	Pagination pagination `json:"pagination"`
}

// Monitor represents a monitor
type Monitor struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes monitorAttributes `json:"attributes,omitempty"`
}

type monitorAttributes struct {
	URL                 string          `json:"url"`
	PronounceableName   string          `json:"pronounceable_name"`
	MonitorType         string          `json:"monitor_type"`
	MonitorGroupID      int             `json:"monitor_group_id"`
	LastCheckedAt       string          `json:"last_checked_at"`
	Status              string          `json:"status"`
	RequiredKeyword     string          `json:"required_keyword"`
	VerifySSL           bool            `json:"verify_ssl"`
	CheckFrequency      int             `json:"check_frequency"`
	Call                bool            `json:"call"`
	SMS                 bool            `json:"sms"`
	Email               bool            `json:"email"`
	Push                bool            `json:"push"`
	TeamWait            int             `json:"team_wait"`
	HTTPMethod          string          `json:"http_method"`
	AuthUsername        string          `json:"auth_username"`
	AuthPassword        string          `json:"auth_password"`
	RequestTimeout      int             `json:"request_timeout"`
	RecoveryPeriod      int             `json:"recovery_period"`
	RequestHeaders      []requestHeader `json:"request_headers"`
	RequestBody         string          `json:"request_body"`
	FollowRedirects     bool            `json:"follow_redirects"`
	RememberCookies     bool            `json:"remember_cookies"`
	Paused              bool            `json:"paused"`
	MaintenanceFrom     string          `json:"maintenance_from"`
	MaintenanceTo       string          `json:"maintenance_to"`
	MaintenanceTimezone string          `json:"maintenance_timezone"`
	Relationships       relationships   `json:"relationships"`
	PausedAt            string          `json:"paused_at"`
	CreatedAt           string          `json:"created_at"`
	UpdatedAt           string          `json:"updated_at"`
	SSLExpiration       int             `json:"ssl_expiration"`
	DomainExpiration    int             `json:"domain_expiration"`
	Regions             []string        `json:"regions"`
	Port                string          `json:"port"`
	ConfirmationPeriod  int             `json:"confirmation_period"`
	ExpectedStatusCodes []int           `json:"expected_status_codes"`
}

type requestHeader struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type relationships struct {
	Policy struct {
		Data string `json:"data,omitempty"`
	} `json:"policy,omitempty"`
	OncallUsers struct {
		Data []user `json:"data,omitempty"`
	} `json:"on_call_users,omitempty"`
	Monitor struct {
		Data data `json:"data"`
	} `json:"monitor,omitempty"`
	Owner struct {
		Data data `json:"data,omitempty"`
	} `json:"owner"`
	WebhookIntegration struct {
		Data data `json:"data,omitempty"`
	} `json:"webhook_integration,omitempty"`
	StatusUpdates struct {
		Data []data
	}
}

type data struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// MonitorResponseTime represents a response containing
// a list of response times for a monitor
type MonitorResponseTime struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes responseTimeAttributes `json:"attributes"`
}

type responseTimeAttributes struct {
	Regions []region `json:"regions"`
}

type region struct {
	Region        string         `json:"region"`
	ResponseTimes []responseTime `json:"response_times"`
}

type responseTime struct {
	At           string  `json:"at"`
	ResponseTime float64 `json:"response_time"`
}

// MonitorAvailability represents a response containing monitor availability
type MonitorAvailability struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes availabilityAttributes `json:"attributes"`
}

type availabilityAttributes struct {
	Availability      float64 `json:"availability,omitempty"`
	TotalDownTime     int     `json:"total_down_time,omitempty"` // time in seconds
	NumberOfIncidents int     `json:"number_of_incidents,omitempty"`
	LongestIncident   int     `json:"longest_incident,omitempty"` // time in seconds
	AverageIncident   int     `json:"average_incident,omitempty"` // time in seconds
}

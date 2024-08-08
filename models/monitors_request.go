package models

import (
	"net/url"
	"strconv"
)

// ListMonitorsQuery represents the query parameters to
// make a request for a list of monitors
type ListMonitorsQuery struct {
	URL               string
	PronounceableName string
	PerPage           int
}

func (q ListMonitorsQuery) ToUrlString() string {
	params := url.Values{}
	params.Add("url", q.URL)
	params.Add("pronounceable_name", q.PronounceableName)
	params.Add("per_page", strconv.Itoa(q.PerPage))
	return params.Encode()
}

// MonitorAvailabilityQuery represents a query for
// monitoring availability within a specified time range.
type MonitorAvailabilityQuery struct {
	MonitorID string // Required
	From      string // Optional Start date (e.g., 2021-01-26)
	To        string // Optional End date (e.g., 2021-01-27), requires `From` to be set
}

// MonitorReqBody represents request body to create a monitor
// URL field is required, every other field is optional
type MonitorReqBody struct {
	URL                 string          `json:"url"`                             // Required
	MonitorType         string          `json:"monitor_type,omitempty"`          // defaults to `status` if empty. access valid values with MonitorTypeList
	PronounceableName   string          `json:"pronounceable_name,omitempty"`    // name of monitor
	Email               bool            `json:"email,omitempty"`                 // Send email alerts
	SMS                 bool            `json:"sms,omitempty"`                   // Send sms alerts
	Call                bool            `json:"call,omitempty"`                  // Phone call alerts
	Push                bool            `json:"push,omitempty"`                  // Should we send a push notification to the on-call person?
	CheckFrequency      int             `json:"check_frequency,omitempty"`       // Check frequency (in seconds)
	RequestHeaders      []requestHeader `json:"request_headers,omitempty"`       // The request headers that will be send with the check
	ExpectedStatusCodes []int           `json:"expected_status_codes,omitempty"` // An array of status codes you expect to receive from your website. These status codes are considered only if the monitor_type is expected_status_code.
	DomainExpiration    int             `json:"domain_expiration,omitempty"`     // How many days before the domain expires do you want to be alerted? Valid values are 1, 2, 3, 7, 14, 30, and 60.
	SSLExpiration       int             `json:"ssl_expiration,omitempty"`        // How many days before the SSL certificate expires do you want to be alerted? Valid values are 1, 2, 3, 7, 14, 30, and 60.
	PolicyID            int             `json:"policy_id,omitempty"`             // Set the escalation policy for the monitor.
	FollowRedirects     bool            `json:"follow_redirects,omitempty"`      // Should we automatically follow redirects when sending the HTTP request?
	TeamWait            int             `json:"team_wait,omitempty"`             // How long to wait before escalating the incident alert to the team. Leave blank to disable escalating to the entire team. In seconds.
	RequiredKeyword     string          `json:"required_keyword,omitempty"`      // Required if monitor_type is set to keyword or udp. We will create a new incident if this keyword is missing on your page.
	Paused              bool            `json:"paused,omitempty"`                // Set to true to pause monitoring
	Port                string          `json:"port,omitempty"`                  // Required if monitor_type is set to tcp, udp, smtp, pop, or imap. tcp and udp monitors accept any ports, while smtp, pop, and imap accept only the specified ports corresponding with their servers (e.g. 25,465,587 for smtp).
	Regions             []string        `json:"regions,omitempty"`               // An array of regions to set. Allowed values are ['us', 'eu', 'as', 'au'] or any subset of these regions.
	MonitorGroupID      string          `json:"monitor_group_id,omitempty"`      // Set this attribute if you want to add this monitor to a monitor group.
	RecoveryPeriod      int             `json:"recovery_period,omitempty"`       // How long the monitor must be up to automatically mark an incident as resolved after being down. In seconds.
	VerifySSL           bool            `json:"verify_ssl,omitempty"`            // verify SSL certificate validity?
	ConfirmationPeriod  int             `json:"confirmation_period,omitempty"`   // Wait time after observing a failure before starting a new incident. In seconds.
	HTTPMethod          string          `json:"http_method,omitempty"`           // HTTP Method used to make a request. Valid options: GET, HEAD, POST, PUT, PATCH
	RequestTimeout      int             `json:"request_timeout,omitempty"`       // How long to wait before timing out the request. In seconds.
	RequestBody         string          `json:"request_body,omitempty"`          // parameter1=first_value&parameter2=another_value
	AuthUsername        string          `json:"auth_username,omitempty"`         // Basic HTTP authentication username to include with the request.
	AuthPassword        string          `json:"auth_password,omitempty"`         // Basic HTTP authentication password to include with the request.
	MaintenanceFrom     string          `json:"maintenance_from,omitempty"`      // Start of the maintenance window each day. We won't check your website during this window. Example: '01:00:00'
	MaintenanceTo       string          `json:"maintenance_to,omitempty"`        // End of the maintenance window each day. Example: '03:00:00'
	MaintenanceTimezone string          `json:"maintenance_timezone,omitempty"`  // Defaults to UTC. The accepted values can be found in the Rails TimeZone documentation. https://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html
	RememberCookies     bool            `json:"remember_cookies,omitempty"`      // keep cookies when redirecting
	PlaywrightScript    string          `json:"playwright_script,omitempty"`     // Playwright script to run
	ScenarioName        string          `json:"scenario_name,omitempty"`         // Name of the scenario identifying the playwright script
}

type monitorTypes struct {
	Status             string // check your website for a 2XX HTTP status code
	ExpectedStatusCode string // check if your website returned one of the values in expected_status_codes.
	Keyword            string // check if your website contains the required_keyword.
	KeywordAbsence     string // check if your website doesn't contain the required_keyword.
	Ping               string // ping your host specified in the url parameter
	TCP                string // test a TCP port at your host specified in the url parameter (port is required)
	UDP                string // test a UDP port at your host specified in the url parameter (port and equired_keyword are required)
	SMTP               string // check for a SMTP server at the host specified in the url parameter (port is required, and can be one of these: 25, 465, 587, or a combination of those ports separated by comma)
	POP                string // check for a POP3 server at the host specified in the url parameter (port is required, and can be 110, 995, or both)
	IMAP               string // check for an IMAP server at the host specified in the url parameter (port is required, and can be 143, 993, or both)
}

// MonitorTypeList is a list of valid betterstack monitor types
var MonitorTypeList = monitorTypes{
	Status:             "status",
	ExpectedStatusCode: "expected_status_code",
	Keyword:            "keyword",
	KeywordAbsence:     "keyword_absence",
	Ping:               "ping",
	TCP:                "tcp",
	UDP:                "udp",
	SMTP:               "smtp",
	POP:                "pop",
	IMAP:               "imap",
}

// var validExpirationValues = map[string]int{
// 	"one": 1, "two": 2, "three": 3, "seven": 7, "fourteen": 14, "thirty": 30, "sixty": 60,
// }

// var validHTTPMethods = map[string]string{
// 	"get": "GET",
// 	"post": "POST",
// 	"patch": "PATCH",
//   "put": "PUT",
//  "head": "HEAD",
// }

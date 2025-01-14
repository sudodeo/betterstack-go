package models

// OncallCalendars represents the response structure for on-call calendars.
type OncallCalendars struct {
	Data       []OncallCalendar `json:"data,omitempty"`
	Included   []User           `json:"included,omitempty"`
	Pagination Pagination       `json:"pagination,omitempty"`
}

// OncallCalendar represents the on-call calendar information.
type OncallCalendar struct {
	ID            string           `json:"id,omitempty"`
	Type          string           `json:"type,omitempty"`
	Attributes    OncallAttributes `json:"attributes,omitempty"`
	Relationships Relationships    `json:"relationships,omitempty"`
}

type OncallAttributes struct {
	Name            string `json:"name,omitempty"`
	DefaultCalendar bool   `json:"default_calendar,omitempty"`
}

type User struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Name       string         `json:"name,omitempty"`
	Attributes UserAttributes `json:"attributes,omitempty"`
}

type UserAttributes struct {
	FirstName    string   `json:"first_name,omitempty"`
	LastName     string   `json:"last_name,omitempty"`
	Email        string   `json:"email,omitempty"`
	PhoneNumbers []string `json:"phone_numbers,omitempty"`
}

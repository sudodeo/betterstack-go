package models

// OncallCalendars represents the response structure for on-call calendars.
type OncallCalendars struct {
	Data       []OncallCalendar `json:"data,omitempty"`
	Included   []user           `json:"included,omitempty"`
	Pagination pagination       `json:"pagination,omitempty"`
}

// OncallCalendar represents the on-call calendar information.
type OncallCalendar struct {
	ID            string           `json:"id,omitempty"`
	Type          string           `json:"type,omitempty"`
	Attributes    oncallAttributes `json:"attributes,omitempty"`
	Relationships relationships    `json:"relationships,omitempty"`
}

type oncallAttributes struct {
	Name            string `json:"name,omitempty"`
	DefaultCalendar bool   `json:"default_calendar,omitempty"`
}

type user struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Name       string         `json:"name,omitempty"`
	Attributes userAttributes `json:"attributes,omitempty"`
}

type userAttributes struct {
	FirstName    string   `json:"first_name,omitempty"`
	LastName     string   `json:"last_name,omitempty"`
	Email        string   `json:"email,omitempty"`
	PhoneNumbers []string `json:"phone_numbers,omitempty"`
}

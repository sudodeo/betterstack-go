package uptime

import (
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListOncallCalendars lists all on-call calendars.
func (bs *Betterstack) ListOncallCalendars() (*models.OncallCalendars, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/on-calls", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.OncallCalendars{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetOncallCalendar returns a single on-call calendar by ID.
func (bs *Betterstack) GetOncallCalendar(calendarID string) (*models.OncallCalendar, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/on-calls/"+calendarID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.OncallCalendar `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// TODO: how do you create calendars? missing in docs

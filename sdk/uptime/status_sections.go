package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

// ListStatusPageSections returns a list of the sections of your status page.
func (bs *Betterstack) ListStatusPageSections(statusPageID string) (*uptime.StatusPageSections, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/status-pages/"+statusPageID+"/sections", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &uptime.StatusPageSections{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetStatusPageSection returns a single statusPage by ID.
func (bs *Betterstack) GetStatusPageSection(statusPageID, sectionID string) (*uptime.StatusPageSection, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/status-pages/"+statusPageID+"/sections/"+sectionID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.StatusPageSection `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateStatusPageSection returns either a newly created statusPage, or validation errors.
func (bs *Betterstack) CreateStatusPageSection(statusPageID string, bodyParams uptime.StatusPageSectionReqBody) (*uptime.StatusPageSection, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/status-pages/"+statusPageID+"/sections/", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.StatusPageSection `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateStatusPageSection update an existing statusPage configuration. Send only the parameters you wish to change (e.g. name)
func (bs *Betterstack) UpdateStatusPageSection(statusPageID, statusPageSectionID string, bodyParams uptime.StatusPageSectionReqBody) (*uptime.StatusPageSection, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/status-pages/"+statusPageID+"/sections/"+statusPageSectionID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.StatusPageSection `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteStatusPageSection permanently deletes an existing statusPage.
func (bs *Betterstack) DeleteStatusPageSection(statusPageID, sectionID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/status-pages/"+statusPageID+"/sections/"+sectionID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

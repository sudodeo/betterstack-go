package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListStatusPages returns a list of all your statusPages.
func (bs *Betterstack) ListStatusPages() (*models.StatusPages, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/status-pages", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.StatusPages{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetStatusPage returns a single statusPage by ID.
func (bs *Betterstack) GetStatusPage(statusPageID string) (*models.StatusPage, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/status-pages/"+statusPageID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.StatusPage `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateStatusPage returns either a newly created statusPage, or validation errors.
func (bs *Betterstack) CreateStatusPage(bodyParams models.StatusPageReqBody) (*models.StatusPage, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/status-pages", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.StatusPage `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateStatusPage update an existing statusPage configuration. Send only the parameters you wish to change (e.g. name)
func (bs *Betterstack) UpdateStatusPage(statusPageID string, bodyParams models.StatusPageReqBody) (*models.StatusPage, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/status-pages/"+statusPageID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.StatusPage `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteStatusPage permanently deletes an existing statusPage.
func (bs *Betterstack) DeleteStatusPage(statusPageID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/status-pages/"+statusPageID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListStatusUpdates returns a list of the reports of your status page.
func (bs *Betterstack) ListStatusUpdates(statusPageID, statusReportID string) (*models.StatusUpdates, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/status-pages/"+statusPageID+"/status-reports/"+statusReportID+"/status-updates", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.StatusUpdates{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetStatusUpdate returns a single statusPage by ID.
func (bs *Betterstack) GetStatusUpdate(statusPageID, statusReportID, statusUpdateID string) (*models.StatusUpdate, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/status-pages/"+statusPageID+"/status-reports/"+statusReportID+"/status-updates/"+statusUpdateID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.StatusUpdate `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateStatusUpdate returns either a newly created statusPage, or validation errors.
func (bs *Betterstack) CreateStatusUpdate(statusPageID, statusReportID string, bodyParams models.StatusUpdateReqBody) (*models.StatusUpdate, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/status-pages/"+statusPageID+"/status-reports/"+statusReportID+"/status-updates", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.StatusUpdate `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateStatusUpdate update an existing statusPage configuration. Send only the parameters you wish to change (e.g. name)
func (bs *Betterstack) UpdateStatusUpdate(statusPageID, statusReportID, statusUpdateID string, bodyParams models.StatusUpdateReqBody) (*models.StatusUpdate, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/status-pages/"+statusPageID+"/status-reports/"+statusReportID+"/status-updates/"+statusUpdateID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.StatusUpdate `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteStatusUpdate permanently deletes an existing statusPage.
func (bs *Betterstack) DeleteStatusUpdate(statusPageID, statusReportID, statusUpdateID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/status-pages/"+statusPageID+"/status-reports/"+statusReportID+"/status-updates/"+statusUpdateID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

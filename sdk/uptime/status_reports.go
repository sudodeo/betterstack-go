package uptime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListStatusPageReports returns a list of the reports of your status page.
func (bs *Betterstack) ListStatusPageReports(statusPageID string) (*models.StatusPageReports, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/status-pages/"+statusPageID+"/status-reports/", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.StatusPageReports{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetStatusPageReport returns a single statusPage by ID.
func (bs *Betterstack) GetStatusPageReport(statusPageID, reportID string) (*models.StatusPageReport, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/status-pages/"+statusPageID+"/status-reports/"+reportID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	resJSON := &models.StatusPageReport{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return resJSON, nil
}

// CreateStatusPageReport returns either a newly created statusPage, or validation errors.
func (bs *Betterstack) CreateStatusPageReport(statusPageID string, bodyParams models.StatusPageReportReqBody) (*models.StatusPageReport, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/status-pages/"+statusPageID+"/status-reports", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	resJSON := &models.StatusPageReport{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return resJSON, nil
}

// UpdateStatusPageReport update an existing statusPage configuration. Send only the parameters you wish to change (e.g. name)
func (bs *Betterstack) UpdateStatusPageReport(statusPageID, statusPageReportID string, bodyParams models.StatusPageReportReqBody) (*models.StatusPageReport, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/status-pages/"+statusPageID+"/status-reports/"+statusPageReportID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	resJSON := &models.StatusPageReport{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return resJSON, nil
}

// DeleteStatusPageReport permanently deletes an existing statusPage.
func (bs *Betterstack) DeleteStatusPageReport(statusPageID, reportID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/status-pages/"+statusPageID+"/status-reports/"+reportID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

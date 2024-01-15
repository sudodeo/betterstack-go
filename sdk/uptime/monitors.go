package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListMonitors returns list of all your monitors
func (bs *Betterstack) ListMonitors(queryParams models.ListMonitorsQuery) (*models.Monitors, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/monitors", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.Monitors{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetMonitor returns a single monitor.
func (bs *Betterstack) GetMonitor(monitorID string) (*models.Monitor, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/monitors/"+monitorID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.Monitor `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// GetResponseTimes returns the response times for a monitor (last 24h).
func (bs *Betterstack) GetResponseTimes(monitorID string) (*models.MonitorResponseTime, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/monitors/"+monitorID+"/response-times", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.MonitorResponseTime `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// GetAvailability returns availability summary for a specific monitor.
func (bs *Betterstack) GetAvailability(queryParams models.MonitorAvailabilityQuery) (*models.MonitorAvailability, error) {
	URL := "/api/v2/monitors/" + queryParams.MonitorID + "/sla"
	if queryParams.From != "" {
		URL = URL + "?from=" + queryParams.From
		if queryParams.To != "" {
			URL = URL + "&to=" + queryParams.To
		}
	}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.MonitorAvailability `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateMonitor returns a newly created monitor or validation errors.
func (bs *Betterstack) CreateMonitor(bodyParams models.MonitorReqBody) (*models.Monitor, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/monitors", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.Monitor `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateMonitor updates existing monitor configuration. Send only the parameters you wish to change (eg. url)
func (bs *Betterstack) UpdateMonitor(monitorID string, bodyParams models.MonitorReqBody) (*models.Monitor, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/monitors/"+monitorID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.Monitor `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteMonitor permanently deletes an existing monitor.
func (bs *Betterstack) DeleteMonitor(monitorID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/monitors/"+monitorID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

// ListMonitorGroups returns a list of all your monitor groups.
func (bs *Betterstack) ListMonitorGroups() (*uptime.Groups, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/monitor-groups", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &uptime.Groups{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetMonitorGroup retrieves a monitor group by its ID
func (bs *Betterstack) GetMonitorGroup(monitorGroupID string) (*uptime.Group, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/monitor-groups/"+monitorGroupID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.Group `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// GetAllMonitorsInGroup returns monitors belonging to the given monitor group.
func (bs *Betterstack) GetAllMonitorsInGroup(monitorGroupID string) (*uptime.Monitors, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/monitor-groups/"+monitorGroupID+"/monitors", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &uptime.Monitors{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateMonitorGroup returns either a newly created monitor group, or validation errors.
func (bs *Betterstack) CreateMonitorGroup(bodyParams uptime.GroupReqBody) (*uptime.Group, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/monitor-groups", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.Group `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateMonitorGroup updates the attributes of an existing monitor group. Send only the parameters you wish to change (e.g. name).
func (bs *Betterstack) UpdateMonitorGroup(monitorGroupID string, bodyParams uptime.GroupReqBody) (*uptime.Group, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/monitor-groups/"+monitorGroupID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.Group `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteMonitorGroup permanently deletes an existing monitor group.
func (bs *Betterstack) DeleteMonitorGroup(monitorGroupID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/monitor-groups/"+monitorGroupID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

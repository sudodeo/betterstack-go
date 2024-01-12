package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

// ListHeartbeatGroups returns a list of all your heartbeat groups.
func (bs *Betterstack) ListHeartbeatGroups() (*uptime.Groups, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/heartbeat-groups", nil)
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

// GetHeartbeatGroup returns an existing heartbeat group by ID.
func (bs *Betterstack) GetHeartbeatGroup(heartbeatID string) (*uptime.Group, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/heartbeat-groups/"+heartbeatID, nil)
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

// CreateHeartbeatGroup returns either a newly created heartbeat group, or validation errors.
func (bs *Betterstack) CreateHeartbeatGroup(bodyParams uptime.GroupReqBody) (*uptime.Group, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/heartbeat-groups", bytes.NewReader(requestBody))
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

// UpdateHeartbeatGroup updates the attributes of an existing heartbeat group. Send only the parameters you wish to change (e.g. name)
func (bs *Betterstack) UpdateHeartbeatGroup(heartbeatID string, bodyParams uptime.GroupReqBody) (*uptime.Group, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/heartbeat-groups/"+heartbeatID, bytes.NewReader(requestBody))
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

// DeleteHeartbeatGroup permanently deletes an existing heartbeat group.
func (bs *Betterstack) DeleteHeartbeatGroup(heartbeatID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/heartbeat-groups/"+heartbeatID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

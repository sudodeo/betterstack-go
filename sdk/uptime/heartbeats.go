package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListHeartbeats returns a list of all your heartbeats.
func (bs *Betterstack) ListHeartbeats(perPage *int) (*models.Heartbeats, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/heartbeats", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.Heartbeats{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetHeartbeat returns a single heartbeat by ID.
func (bs *Betterstack) GetHeartbeat(heartbeatID string) (*models.Heartbeat, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/heartbeats/"+heartbeatID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.Heartbeat `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateHeartbeat returns either a newly created heartbeat, or validation errors.
func (bs *Betterstack) CreateHeartbeat(bodyParams models.HeartbeatReqBody) (*models.Heartbeat, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/heartbeats", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.Heartbeat `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateHeartbeat update an existing heartbeat configuration. Send only the parameters you wish to change (e.g. name)
func (bs *Betterstack) UpdateHeartbeat(heartbeatID string, bodyParams models.HeartbeatReqBody) (*models.Heartbeat, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/heartbeats/"+heartbeatID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.Heartbeat `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteHeartbeat permanently deletes an existing heartbeat.
func (bs *Betterstack) DeleteHeartbeat(heartbeatID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/heartbeats/"+heartbeatID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

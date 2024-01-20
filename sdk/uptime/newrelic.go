package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListNewRelics returns a list of all your New Relic integrations.
func (bs *Betterstack) ListNewRelics() (*models.NewRelics, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/new-relic-integrations", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.NewRelics{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetNewRelic returns a single New Relic integration by ID.
func (bs *Betterstack) GetNewRelic(newRelicID string) (*models.NewRelic, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/new-relic-integrations/"+newRelicID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.NewRelic `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateNewRelic returns either a newly created New Relic integration, or validation errors.
func (bs *Betterstack) CreateNewRelic(bodyParams models.NewRelicReqBody) (*models.NewRelic, error) {
	// TODO:  figure out required body params
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/new-relic-integrations", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.NewRelic `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateNewRelic updates an existing New Relic integration. Send only the parameters you wish to change (eg. paused).
func (bs *Betterstack) UpdateNewRelic(newRelicID string, bodyParams models.NewRelicReqBody) (*models.NewRelic, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/new-relic-integrations/"+newRelicID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.NewRelic `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteNewRelic permanently deletes an existing New Relic integration.
func (bs *Betterstack) DeleteNewRelic(newRelicID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/new-relic-integrations/"+newRelicID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

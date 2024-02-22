package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models"
)

// ListEscalationPolicies returns list of all your escalation policies
func (bs *Betterstack) ListEscalationPolicies() (*models.EscalationPolicies, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/policies", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.EscalationPolicies{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetEscalationPolicy returns a single escalation policy.
func (bs *Betterstack) GetEscalationPolicy(policyID string) (*models.EscalationPolicy, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/policies/"+policyID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.EscalationPolicy `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateEscalationPolicy returns a newly created escalation policy or validation errors.
func (bs *Betterstack) CreateEscalationPolicy(bodyParams models.PolicyReqBody) (*models.EscalationPolicy, error) {
	// TODO:  figure out required body params
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/policies", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.EscalationPolicy `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateEscalationPolicy updates existing escalation policy configuration. Send only the parameters you wish to change (eg. url)
func (bs *Betterstack) UpdateEscalationPolicy(policyID string, bodyParams models.PolicyReqBody) (*models.EscalationPolicy, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/policies/"+policyID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data models.EscalationPolicy `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteEscalationPolicy permanently deletes an existing escalation policy.
func (bs *Betterstack) DeleteEscalationPolicy(policyID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/policies/"+policyID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

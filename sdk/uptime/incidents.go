package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

// ListIncidents returns a list of all incidents with the option to filter incidents by monitor.
func (bs *Betterstack) ListIncidents() (*uptime.Incidents, error) {
	// TODO: add query params
	req, err := http.NewRequest(http.MethodGet, "/api/v2/incidents", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &uptime.Incidents{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetIncident returns a single incident.
func (bs *Betterstack) GetIncident(incidentID string) (*uptime.Incident, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/incidents/"+incidentID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.Incident `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// GetIncidentTimelineEvents returns a list of timeline items for the given incident.
func (bs *Betterstack) GetIncidentTimelineEvents(incidentID string) (*uptime.IncidentTimelineEvents, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/incidents/"+incidentID+"/timeline", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &uptime.IncidentTimelineEvents{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateIncident returns a newly created incident or validation errors.
func (bs *Betterstack) CreateIncident(bodyParams uptime.IncidentReqBody) (*uptime.Incident, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/incidents", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.Incident `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// AcknowledgeIncident acknowledges an incident.
func (bs *Betterstack) AcknowledgeIncident(incidentID, email string) (*uptime.Incident, error) {
	type reqBody struct {
		AcknowledgedBy string `json:"acknowledged_by"`
	}
	bodyParams := reqBody{
		AcknowledgedBy: email,
	}
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/incidents/"+incidentID+"/acknowledge", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.Incident `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// ResolveIncident resolves an incident.
func (bs *Betterstack) ResolveIncident(incidentID, email string) (*uptime.Incident, error) {
	type reqBody struct {
		ResolvedBy string `json:"resolved_by"`
	}
	bodyParams := reqBody{
		ResolvedBy: email,
	}
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/incidents/"+incidentID+"/resolve", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.Incident `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteIncident permanently deletes an existing incident.
func (bs *Betterstack) DeleteIncident(incidentID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/incidents/"+incidentID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

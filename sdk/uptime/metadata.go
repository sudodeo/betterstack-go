package uptime

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

// ListMetadata returns list of all metadata
func (bs *Betterstack) ListMetadata() (*uptime.Metadata, error) {
	// TODO: query params
	req, err := http.NewRequest(http.MethodGet, "/api/v2/metadata", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &uptime.Metadata{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetMetadataRecord returns a single metadata record.
func (bs *Betterstack) GetMetadataRecord(metadataID string) (*uptime.MetadataRecord, error) {
	req, err := http.NewRequest(http.MethodGet, "/api/v2/metadata/"+metadataID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.MetadataRecord `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// CreateMetadataRecord creates a new metadata record or returns validation errors.
func (bs *Betterstack) CreateMetadataRecord(bodyParams uptime.MetadataRecordReqBody) (*uptime.MetadataRecord, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/metadata", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.MetadataRecord `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// UpdateMetadataRecord updates an existing metadata record configuration. Send only the parameters you wish to change (e.g., URL).
func (bs *Betterstack) UpdateMetadataRecord(metadataID string, bodyParams uptime.MetadataRecordReqBody) (*uptime.MetadataRecord, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, "/api/v2/metadata/"+metadataID, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data uptime.MetadataRecord `json:"data"`
	}

	resJSON := &response{}
	err = json.Unmarshal(body, resJSON)
	if err != nil {
		return nil, err
	}

	return &resJSON.Data, nil
}

// DeleteMetadataRecord permanently deletes an existing metadata record.
func (bs *Betterstack) DeleteMetadataRecord(metadataID string) error {
	req, err := http.NewRequest(http.MethodDelete, "/api/v2/metadata/"+metadataID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

package logs

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/sudodeo/betterstack-go/models"
)

/*
ListSources Returns a list of your team's existing sources.
*/
func (bs *Betterstack) ListSources(page, perPage *int) (*models.Sources, error) {
	if perPage != nil && *perPage > 50 {
		return nil, errors.New("perPage can only be max 50")
	} else if perPage != nil && *perPage < 1 {
		return nil, errors.New("perPage must be greater than 0")
	}

	if *page < 1 {
		return nil, errors.New("page must be greater than 0")
	}

	queryParams := "?page=" + strconv.Itoa(*page)

	if perPage != nil {
		queryParams = queryParams + "&per_page=" + strconv.Itoa(*perPage)
	}

	req, err := http.NewRequest("GET", "/api/v1/sources"+queryParams, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.Sources{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetSource retrieves the source data for a given source ID.
func (bs *Betterstack) GetSource(sourceID string) (*models.Source, error) {
	req, err := http.NewRequest("GET", "/api/v1/sources/"+sourceID, nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &models.Source{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateSource creates a new source in the Betterstack SDK.
// It takes a bodyParams parameter of type models.CreateSourceBodyParams, which contains the necessary information to create the source.
// It returns a *models.Source and an error.
func (bs *Betterstack) CreateSource(bodyParams models.CreateSourceBodyParams) (*models.Source, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "/api/v1/sources/", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data *models.Source `json:"data"`
	}
	bodyJSON := &response{}
	err = json.Unmarshal(body, bodyJSON)
	if err != nil {
		return nil, err
	}

	return bodyJSON.Data, nil
}

// UpdateSource Updates an existing source. Send only the parameters you wish to change (e.g. name )
func (bs *Betterstack) UpdateSource(sourceID string, bodyParams models.UpdateSourceBodyParams) (*models.Source, error) {
	requestBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", "/api/v1/sources/"+sourceID, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	type response struct {
		Data *models.Source `json:"data"`
	}
	bodyJSON := &response{}
	err = json.Unmarshal(body, bodyJSON)
	if err != nil {
		return nil, err
	}

	return bodyJSON.Data, nil
}

// DeleteSource deletes a source with the specified sourceID from the Betterstack SDK.
// It sends a DELETE request to the "/api/v1/sources/{sourceID}" endpoint.
// If the request is successful, it returns nil. Otherwise, it returns an error.
func (bs *Betterstack) DeleteSource(sourceID string) error {
	req, err := http.NewRequest("DELETE", "/api/v1/sources/"+sourceID, nil)
	if err != nil {
		return err
	}

	_, err = bs.MakeAPIRequest(req)
	if err != nil {
		return err
	}

	return nil
}

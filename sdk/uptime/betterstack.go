package uptime

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// Betterstack represents the client for interacting with the Betterstack API.
type Betterstack struct {
	client  *http.Client
	baseURL string
	Token   string
}

// FailedRequest represents a failed request in the BetterStack SDK.
type FailedRequest struct {
	Errors        string `json:"errors"`
	InvalidValues struct {
		SourceIDs string `json:"source_ids"`
		Batch     string `json:"batch"`
		From      string `json:"from"`
		To        string `json:"to"`
		Order     string `json:"order"`
	} `json:"invalid_values,omitempty"`
}

// MakeAPIRequest sends an API request using the provided http.Request object and returns the response body as a byte array.
// If the request fails or the response status code is not 200, an error is returned.
func (bs *Betterstack) MakeAPIRequest(req *http.Request) ([]byte, error) {
	var errData *FailedRequest
	req.Header = http.Header{
		"accept":        {"application/json"},
		"content-type":  {"application/json"},
		"Authorization": {"Bearer " + bs.Token},
	}

	req.URL = &url.URL{
		Scheme: "https",
		Host:   bs.baseURL,
		Path:   req.URL.Path,
	}

	res, err := bs.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Success is indicated with 2xx status codes:
	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		if err = json.Unmarshal(body, &errData); err != nil {
			return nil, err
		}

		return nil, &BetterStackError{
			StatusCode: res.StatusCode,
			Message:    errData.Errors,
		}
	}
	if res.StatusCode == 204 {
		return nil, nil
	}

	return body, nil
}

// NewFromENV creates a new instance of the betterstack struct using the UPTIME_API_TOKEN environment variable.
// It returns a pointer to the betterstack struct and an error if any.
func NewFromENV() (*Betterstack, error) {
	// only way to load env variables in test
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("unable to identify current directory")
	}

	basepath := filepath.Dir(file)
	envPath := filepath.Join(basepath, "../../.env")
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	uptimeAPIToken := os.Getenv("UPTIME_API_TOKEN")

	bs := Betterstack{
		Token:   uptimeAPIToken,
		client:  &http.Client{},
		baseURL: "uptime.betterstack.com",
	}

	return &bs, nil
}

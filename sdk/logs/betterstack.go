package logs

import (
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
	token   string
}

// MakeAPIRequest sends an API request using the provided http.Request object and returns the response body as a byte array.
// If the request fails or the response status code is not 200, an error is returned.
func (bs *Betterstack) MakeAPIRequest(req *http.Request) ([]byte, error) {
	req.Header = http.Header{
		"accept":        {"application/json"},
		"content-type":  {"application/json"},
		"Authorization": {"Bearer " + bs.token},
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
		return nil, &BetterStackError{
			StatusCode: res.StatusCode,
			Message:    string(body),
		}
	}
	if res.StatusCode == 204 {
		return nil, nil
	}

	return body, nil
}

// NewFromENV creates a new instance of the betterstack struct using the LOGS_API_TOKEN environment variable.
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

	logsAPIToken := os.Getenv("LOGS_API_TOKEN")

	bs := Betterstack{
		token:   logsAPIToken,
		client:  &http.Client{},
		baseURL: "logs.betterstack.com",
	}

	return &bs, nil
}

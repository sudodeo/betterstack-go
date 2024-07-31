package uptime

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/sudodeo/betterstack-go/sdk"
)

// Betterstack represents the client for interacting with the Betterstack API.
type Betterstack struct {
	client  *http.Client
	baseURL string
	Token   string
}

// MakeAPIRequest sends an API request using the provided http.Request object and returns the response body as a byte array.
// If the request fails or the response status code is not 200, an error is returned.
func (bs *Betterstack) MakeAPIRequest(req *http.Request, rawQueryStrings ...string) ([]byte, error) {
	req.Header = http.Header{
		"accept":        {"application/json"},
		"content-type":  {"application/json"},
		"Authorization": {"Bearer " + bs.Token},
	}

	var rawQueryString string
	if len(rawQueryStrings) > 0 {
		for _, rawQuery := range rawQueryStrings {
			rawQueryString += "&" + rawQuery
		}
	}

	req.URL = &url.URL{
		Scheme:   "https",
		Host:     bs.baseURL,
		Path:     req.URL.Path,
		RawQuery: rawQueryString,
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

// NewFromENV creates a new instance of the Betterstack struct using the UPTIME_API_TOKEN environment variable.
// It returns a pointer to the Betterstack struct and an error if any.
func NewFromENV() (*Betterstack, error) {
	if os.Getenv("env") == "test" {
		// Only load environment variables from a specific path during testing
		err := loadEnvForTest()
		if err != nil {
			return nil, fmt.Errorf("failed to load environment variables for test: %v", err)
		}
	}

	uptimeAPIToken := os.Getenv("UPTIME_API_TOKEN")

	bs := Betterstack{
		Token:   uptimeAPIToken,
		client:  &http.Client{},
		baseURL: "uptime.betterstack.com",
	}

	return &bs, nil
}

func loadEnvForTest() error {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to identify current directory")
	}

	basepath := filepath.Dir(file)
	envPath := filepath.Join(basepath, "../../.env")

	err := godotenv.Load(envPath)
	if err != nil {
		return fmt.Errorf("failed to load .env file: %v", err)
	}

	return nil
}

// New creates a new instance of the Betterstack struct using the provided config.
func New(config *sdk.Config) (*Betterstack, error) {
	if config == nil {
		return nil, fmt.Errorf("config is required")
	}

	if config.APIType != "uptime" {
		return nil, fmt.Errorf("invalid API type: %s", config.APIType)
	}

	uptimeAPIToken := config.APIToken

	bs := Betterstack{
		Token:   uptimeAPIToken,
		client:  &http.Client{},
		baseURL: "uptime.betterstack.com",
	}

	return &bs, nil
}

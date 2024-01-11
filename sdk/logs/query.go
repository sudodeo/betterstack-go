package logs

import (
	"encoding/json"
	"net/http"

	"github.com/sudodeo/betterstack-go/models/logs"
)

// FetchLogs fetches logs from the Betterstack API.
func (bs *Betterstack) FetchLogs() (*logs.Logs, error) {
	req, err := http.NewRequest("GET", "/api/v1/query", nil)
	if err != nil {
		return nil, err
	}

	body, err := bs.MakeAPIRequest(req)
	if err != nil {
		return nil, err
	}

	data := &logs.Logs{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

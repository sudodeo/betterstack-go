package logs

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sudodeo/betterstack-go/models/logs"
)

// FetchLogs fetches logs from the Betterstack API.
func (bs *Betterstack) FetchLogs(queryParams *logs.FetchLogsParams) (*logs.Logs, error) {
	var reqBody io.Reader

	if queryParams != nil {
		reqBodyBytes, err := json.Marshal(queryParams)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(reqBodyBytes)
	}

	req, err := http.NewRequest(http.MethodGet, "/api/v1/query", reqBody)
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

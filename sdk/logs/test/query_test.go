package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
)

func TestFetchLogs(t *testing.T) {
	// Test case 1: Successful response without query params
	_logs, err := bs.FetchLogs(nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, _logs)
	assert.IsType(t, &models.Logs{}, _logs)

	// Test case 2: Successful response with query params
	querParams := &models.FetchLogsParams{
		Batch: 50,
	}
	_logs, err = bs.FetchLogs(querParams)
	assert.Nil(t, err)
	assert.NotEmpty(t, _logs)
	assert.IsType(t, &models.Logs{}, _logs)
}

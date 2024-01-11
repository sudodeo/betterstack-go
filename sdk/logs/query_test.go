package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models/logs"
)

func TestFetchLogs(t *testing.T) {
	_logs, err := bs.FetchLogs()
	assert.Nil(t, err)
	assert.NotEmpty(t, _logs)

	assert.IsType(t, &logs.Logs{}, _logs)
}

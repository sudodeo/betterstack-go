package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

var heartbeatID string

func TestListHeartbeats(t *testing.T) {
	// Test case 1: Successful API response without perPage argument
	heartbeats, err := bs.ListHeartbeats(nil)
	assert.Nil(t, err)
	assert.NotNil(t, heartbeats)
	assert.IsType(t, &uptime.Heartbeats{}, heartbeats)

	// Test case 2: Successful API response with perPage argument
	perPage := 2
	heartbeats, err = bs.ListHeartbeats(&perPage)
	assert.Nil(t, err)
	assert.NotNil(t, heartbeats)
	assert.IsType(t, &uptime.Heartbeats{}, heartbeats)
}

func TestCreateHeartbeat(t *testing.T) {
	reqBody := uptime.HeartbeatReqBody{
		Name:   "heartbeat_test",
		Period: 30,
		Grace:  9,
	}
	heartbeat, err := bs.CreateHeartbeat(reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, heartbeat)
	assert.IsType(t, &uptime.Heartbeat{}, heartbeat)
	assert.Equal(t, reqBody.Name, heartbeat.Attributes.Name)
	assert.Equal(t, reqBody.Period, heartbeat.Attributes.Period)
	assert.Equal(t, reqBody.Grace, heartbeat.Attributes.Grace)
	heartbeatID = heartbeat.ID
}

func TestGetHeartbeat(t *testing.T) {
	heartbeat, err := bs.GetHeartbeat(heartbeatID)
	assert.Nil(t, err)
	assert.NotNil(t, heartbeat)
	assert.IsType(t, &uptime.Heartbeat{}, heartbeat)
	assert.Equal(t, heartbeatID, heartbeat.ID)
}

func TestUpdateHeartbeat(t *testing.T) {
	reqBody := uptime.HeartbeatReqBody{
		Name:   "heartbeat_test_update",
		Period: 50,
		Grace:  15,
	}
	heartbeat, err := bs.UpdateHeartbeat(heartbeatID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, heartbeat)
	assert.IsType(t, &uptime.Heartbeat{}, heartbeat)
	assert.Equal(t, heartbeatID, heartbeat.ID)
	assert.Equal(t, reqBody.Name, heartbeat.Attributes.Name)
	assert.Equal(t, reqBody.Period, heartbeat.Attributes.Period)
	assert.Equal(t, reqBody.Grace, heartbeat.Attributes.Grace)
}

func TestDeleteHeartbeat(t *testing.T) {
	err := bs.DeleteHeartbeat(heartbeatID)
	assert.Nil(t, err)
}

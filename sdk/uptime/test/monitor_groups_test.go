package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
)

var monitorGroupID string

func TestListMonitorGroups(t *testing.T) {
	monitorGroups, err := bs.ListMonitorGroups()
	assert.Nil(t, err)
	assert.NotNil(t, monitorGroups)
	assert.IsType(t, &models.Groups{}, monitorGroups)
}

func TestCreateMonitorGroup(t *testing.T) {
	reqBody := models.GroupReqBody{
		Name: "test2_group",
	}
	monitorGroup, err := bs.CreateMonitorGroup(reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, monitorGroup)
	assert.IsType(t, &models.Group{}, monitorGroup)
	monitorGroupID = monitorGroup.ID
}

func TestGetAllMonitorsInGroup(t *testing.T) {
	monitors, err := bs.GetAllMonitorsInGroup(monitorGroupID)
	assert.Nil(t, err)
	assert.NotNil(t, monitors)
	assert.IsType(t, &models.Monitors{}, monitors)
}

func TestGetMonitorGroup(t *testing.T) {
	monitorGroup, err := bs.GetMonitorGroup(monitorGroupID)
	assert.Nil(t, err)
	assert.NotNil(t, monitorGroup)
	assert.IsType(t, &models.Group{}, monitorGroup)
	assert.Equal(t, monitorGroupID, monitorGroup.ID)
}

func TestUpdateMonitorGroup(t *testing.T) {
	reqBody := models.GroupReqBody{
		Name:   "update test group",
		Paused: true,
	}
	monitorGroup, err := bs.UpdateMonitorGroup(monitorGroupID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, monitorGroup)
	assert.IsType(t, &models.Group{}, monitorGroup)
	assert.Equal(t, monitorGroupID, monitorGroup.ID)
}
func TestDeleteMonitorGroup(t *testing.T) {
	err := bs.DeleteMonitorGroup(monitorGroupID)
	assert.Nil(t, err)
}

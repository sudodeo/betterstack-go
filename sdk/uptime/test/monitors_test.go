package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
)

var monitorID string

func TestListMonitors(t *testing.T) {
	queryParams := models.ListMonitorsQuery{}
	monitors, err := bs.ListMonitors(queryParams)
	assert.Nil(t, err)
	assert.NotNil(t, monitors)
	assert.IsType(t, &models.Monitors{}, monitors)
}

func TestCreateMonitor(t *testing.T) {
	reqBody := models.MonitorReqBody{
		URL: "https://test.com",
	}
	monitor, err := bs.CreateMonitor(reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, monitor)
	assert.IsType(t, &models.Monitor{}, monitor)
	assert.Equal(t, reqBody.URL, monitor.Attributes.URL)
	monitorID = monitor.ID
}

func TestGetMonitor(t *testing.T) {
	monitor, err := bs.GetMonitor(monitorID)
	assert.Nil(t, err)
	assert.NotNil(t, monitor)
	assert.IsType(t, &models.Monitor{}, monitor)
	assert.Equal(t, monitorID, monitor.ID)
}

func TestGetResponseTimes(t *testing.T) {
	monitorResponseTime, err := bs.GetResponseTimes(monitorID)
	assert.Nil(t, err)
	assert.NotNil(t, monitorResponseTime)
	assert.IsType(t, &models.MonitorResponseTime{}, monitorResponseTime)
	assert.Equal(t, monitorID, monitorResponseTime.ID)
}

func TestGetAvailability(t *testing.T) {
	// Test case 1: Successful API response with only monitorID
	queryParams := models.MonitorAvailabilityQuery{
		MonitorID: monitorID,
	}
	monitorAvailability, err := bs.GetAvailability(queryParams)
	assert.Nil(t, err)
	assert.NotNil(t, monitorAvailability)
	assert.IsType(t, &models.MonitorAvailability{}, monitorAvailability)
	assert.Equal(t, monitorID, monitorAvailability.ID)

	// Test case 2: Successful API response with monitorID, and time set
	queryParams = models.MonitorAvailabilityQuery{
		MonitorID: monitorID,
		From:      "2023-12-26",
		To:        "2023-12-31",
	}
	monitorAvailability, err = bs.GetAvailability(queryParams)
	assert.Nil(t, err)
	assert.NotNil(t, monitorAvailability)
	assert.IsType(t, &models.MonitorAvailability{}, monitorAvailability)
	assert.Equal(t, monitorID, monitorAvailability.ID)

	// Test case 3: Failed API response due to missing monitorID
	queryParams = models.MonitorAvailabilityQuery{}
	monitorAvailability, err = bs.GetAvailability(queryParams)
	assert.NotNil(t, err)
	assert.Nil(t, monitorAvailability)
}

func TestUpdateMonitor(t *testing.T) {
	reqBody := models.MonitorReqBody{
		URL: "https://test-update.com",
	}
	monitor, err := bs.UpdateMonitor(monitorID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, monitor)
	assert.IsType(t, &models.Monitor{}, monitor)
	assert.Equal(t, monitorID, monitor.ID)
	assert.Equal(t, reqBody.URL, monitor.Attributes.URL)
}

func TestDeleteMonitor(t *testing.T) {
	err := bs.DeleteMonitor(monitorID)
	assert.Nil(t, err)
}

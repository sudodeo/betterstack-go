package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
)

var statusUpdateID, newStatusReportID string

func TestCreateStatusUpdate(t *testing.T) {
	monitor, err := bs.CreateMonitor(
		models.MonitorCreateReqBody{
			URL: "https://google.com",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, monitor)
	monitorID = monitor.ID

	newStatusPage, err := bs.CreateStatusPage(models.StatusPageReqBody{
		CompanyName: "test_company_update",
		Subdomain:   "test-subdomain-section",
		Timezone:    "Casablanca",
	})

	assert.Nil(t, err)
	assert.NotNil(t, newStatusPage)

	newStatusPageID = newStatusPage.ID

	newResourceReqBody := models.StatusPageResourceReqBody{
		ResourceType: "Monitor",
		ResourceID:   monitorID,
		PublicName:   "test_resource",
	}

	statusResource, err := bs.CreateStatusPageResource(newStatusPageID, newResourceReqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusResource)

	newResourceID = statusResource.ID

	newReportReqBody := models.StatusPageReportReqBody{
		Title: "New report",
		AffectedResources: []models.AffectedResource{
			{StatusPageResourceID: newResourceID, Status: "downtime"},
		},
		Message: "status update report",
	}

	statusPageReport, err := bs.CreateStatusPageReport(newStatusPageID, newReportReqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReport)

	newStatusReportID = statusPageReport.Data.ID

	reqBody := models.StatusUpdateReqBody{
		AffectedResources: []models.AffectedResource{
			{StatusPageResourceID: newResourceID, Status: "downtime"},
		},
		Message: "status update report",
	}

	statusUpdate, err := bs.CreateStatusUpdate(newStatusPageID, newStatusReportID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusUpdate)

	assert.IsType(t, &models.StatusUpdate{}, statusUpdate)
	assert.Equal(t, reqBody.Message, statusUpdate.Attributes.Message)
	assert.Equal(t, reqBody.AffectedResources[0].StatusPageResourceID, statusUpdate.Attributes.AffectedResources[0].StatusPageResourceID)
	statusUpdateID = statusUpdate.ID
}

func TestListStatusUpdates(t *testing.T) {
	// Test case 1: Successful API response
	statusUpdates, err := bs.ListStatusUpdates(newStatusPageID, newStatusReportID)
	assert.Nil(t, err)
	assert.NotNil(t, statusUpdates)
	assert.IsType(t, &models.StatusUpdates{}, statusUpdates)
}

func TestGetStatusUpdate(t *testing.T) {
	statusUpdate, err := bs.GetStatusUpdate(newStatusPageID, newStatusReportID, statusUpdateID)
	assert.Nil(t, err)
	assert.NotNil(t, statusUpdate)
	assert.IsType(t, &models.StatusUpdate{}, statusUpdate)
	assert.Equal(t, statusUpdateID, statusUpdate.ID)
}

func TestUpdateStatusUpdate(t *testing.T) {
	reqBody := models.StatusUpdateReqBody{
		Message: "update message",
		AffectedResources: []models.AffectedResource{
			{StatusPageResourceID: newResourceID, Status: "resolved"},
		},
	}

	statusUpdate, err := bs.UpdateStatusUpdate(newStatusPageID, newStatusReportID, statusUpdateID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusUpdate)
	assert.IsType(t, &models.StatusUpdate{}, statusUpdate)
	assert.Equal(t, reqBody.Message, statusUpdate.Attributes.Message)
}

func TestDeleteStatusUpdate(t *testing.T) {
	err := bs.DeleteStatusUpdate(newStatusPageID, newStatusReportID, statusUpdateID)
	assert.Nil(t, err)

	err = bs.DeleteStatusPage(newStatusPageID)
	assert.Nil(t, err)

	err = bs.DeleteMonitor(monitorID)
}

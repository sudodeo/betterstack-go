package uptime_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
)

var (
	statusPageReportID, newResourceID string
)

func TestCreateStatusPageReport(t *testing.T) {
	monitor, err := bs.CreateMonitor(
		models.MonitorReqBody{
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

	statusPageResource, err := bs.CreateStatusPageResource(newStatusPageID, newResourceReqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageResource)

	newResourceID = statusPageResource.ID

	reqBody := models.StatusPageReportReqBody{
		Title: "New report",
		AffectedResources: []models.AffectedResource{
			{StatusPageResourceID: newResourceID, Status: "downtime"},
		},
		Message: "status update report",
	}

	statusPageReport, err := bs.CreateStatusPageReport(newStatusPageID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReport)
	fmt.Println(statusPageReport)
	assert.IsType(t, &models.StatusPageReport{}, statusPageReport)
	assert.Equal(t, reqBody.Title, statusPageReport.Data.Attributes.Title)
	assert.Equal(t, reqBody.Message, statusPageReport.Included[0].Attributes.Message)
	statusPageReportID = statusPageReport.Data.ID
}

func TestListStatusPageReports(t *testing.T) {
	// Test case 1: Successful API response
	statusPageReports, err := bs.ListStatusPageReports(newStatusPageID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReports)
	assert.IsType(t, &models.StatusPageReports{}, statusPageReports)
}

func TestGetStatusPageReport(t *testing.T) {
	statusPageReport, err := bs.GetStatusPageReport(newStatusPageID, statusPageReportID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReport)
	assert.IsType(t, &models.StatusPageReport{}, statusPageReport)
	assert.Equal(t, statusPageReportID, statusPageReport.Data.ID)
}

func TestUpdateStatusPageReport(t *testing.T) {
	reqBody := models.StatusPageReportReqBody{
		Title: "New report update",
	}

	statusPageReport, err := bs.UpdateStatusPageReport(newStatusPageID, statusPageReportID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReport)
	assert.IsType(t, &models.StatusPageReport{}, statusPageReport)
	assert.Equal(t, reqBody.Title, statusPageReport.Data.Attributes.Title)
}

func TestDeleteStatusPageReport(t *testing.T) {
	err := bs.DeleteStatusPageReport(newStatusPageID, statusPageReportID)
	assert.Nil(t, err)

	err = bs.DeleteStatusPage(newStatusPageID)
	assert.Nil(t, err)

	err = bs.DeleteMonitor(newResourceID)
}

package uptime_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

var (
	statusPageReportID, newResourceID string
)

func TestCreateStatusPageReport(t *testing.T) {
	monitor, err := bs.CreateMonitor(
		uptime.MonitorReqBody{
			URL: "https://google.com",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, monitor)
	monitorID = monitor.ID

	newStatusPage, err := bs.CreateStatusPage(uptime.StatusPageReqBody{
		CompanyName: "test_company_update",
		Subdomain:   "test-subdomain-section",
		Timezone:    "Casablanca",
	})

	assert.Nil(t, err)
	assert.NotNil(t, newStatusPage)
	newStatusPageID = newStatusPage.ID

	newResourceReqBody := uptime.StatusPageResourceReqBody{
		ResourceType: "Monitor",
		ResourceID:   monitorID,
		PublicName:   "test_resource",
	}

	statusPageResource, err := bs.CreateStatusPageResource(newStatusPageID, newResourceReqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageResource)

	newResourceID = statusPageResource.ID

	reqBody := uptime.StatusPageReportReqBody{
		Title: "New report",
		AffectedResources: []uptime.AffectedResource{
			{StatusPageResourceID: newResourceID, Status: "downtime"},
		},
		Message: "status update report",
	}

	statusPageReport, err := bs.CreateStatusPageReport(newStatusPageID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReport)
	fmt.Println(statusPageReport)
	assert.IsType(t, &uptime.StatusPageReport{}, statusPageReport)
	assert.Equal(t, reqBody.Title, statusPageReport.Data.Attributes.Title)
	assert.Equal(t, reqBody.Message, statusPageReport.Included[0].Attributes.Message)
	statusPageReportID = statusPageReport.Data.ID
}

func TestListStatusPageReports(t *testing.T) {
	// Test case 1: Successful API response
	statusPageReports, err := bs.ListStatusPageReports(newStatusPageID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReports)
	assert.IsType(t, &uptime.StatusPageReports{}, statusPageReports)
}

func TestGetStatusPageReport(t *testing.T) {
	statusPageReport, err := bs.GetStatusPageReport(newStatusPageID, statusPageReportID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReport)
	assert.IsType(t, &uptime.StatusPageReport{}, statusPageReport)
	assert.Equal(t, statusPageReportID, statusPageReport.Data.ID)
}

func TestUpdateStatusPageReport(t *testing.T) {
	reqBody := uptime.StatusPageReportReqBody{
		Title: "New report update",
	}

	statusPageReport, err := bs.UpdateStatusPageReport(newStatusPageID, statusPageReportID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageReport)
	assert.IsType(t, &uptime.StatusPageReport{}, statusPageReport)
	assert.Equal(t, reqBody.Title, statusPageReport.Data.Attributes.Title)
}

func TestDeleteStatusPageReport(t *testing.T) {
	err := bs.DeleteStatusPageReport(newStatusPageID, statusPageReportID)
	assert.Nil(t, err)

	err = bs.DeleteStatusPage(newStatusPageID)
	assert.Nil(t, err)

	err = bs.DeleteMonitor(newResourceID)
}

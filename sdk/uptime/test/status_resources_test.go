package uptime_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
)

var (
	statusPageResourceID string
)

func TestCreateStatusPageResource(t *testing.T) {
	monitor, err := bs.CreateMonitor(
		models.MonitorCreateReqBody{
			URL: "https://google.com",
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, monitor)
	monitorID = monitor.ID
	reqBody := models.StatusPageResourceReqBody{
		ResourceType: "Monitor",
		ResourceID:   monitorID,
		PublicName:   "test_resource",
	}

	newStatusPage, err := bs.CreateStatusPage(models.StatusPageReqBody{
		CompanyName: "test_company_update",
		Subdomain:   "test-subdomain-section",
		Timezone:    "Casablanca",
	})

	assert.Nil(t, err)
	assert.NotNil(t, newStatusPage)
	newStatusPageID = newStatusPage.ID
	statusPageResource, err := bs.CreateStatusPageResource(newStatusPageID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageResource)
	assert.IsType(t, &models.StatusPageResource{}, statusPageResource)
	assert.Equal(t, reqBody.PublicName, statusPageResource.Attributes.PublicName)
	assert.Equal(t, reqBody.ResourceID, strconv.Itoa(statusPageResource.Attributes.ResourceID))
	statusPageResourceID = statusPageResource.ID
}

func TestListStatusPageResources(t *testing.T) {
	// Test case 1: Successful API response
	statusPageResources, err := bs.ListStatusPageResources(newStatusPageID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageResources)
	assert.IsType(t, &models.StatusPageResources{}, statusPageResources)
}

func TestGetStatusPageResource(t *testing.T) {
	statusPageResource, err := bs.GetStatusPageResource(newStatusPageID, statusPageResourceID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageResource)
	assert.IsType(t, &models.StatusPageResource{}, statusPageResource)
	assert.Equal(t, statusPageResourceID, statusPageResource.ID)
}

func TestUpdateStatusPageResource(t *testing.T) {
	reqBody := models.StatusPageResourceReqBody{
		PublicName:   "test_resource_update",
		ResourceType: "Monitor",
		ResourceID:   monitorID,
	}
	statusPageResource, err := bs.UpdateStatusPageResource(newStatusPageID, statusPageResourceID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageResource)
	assert.IsType(t, &models.StatusPageResource{}, statusPageResource)
	assert.Equal(t, reqBody.PublicName, statusPageResource.Attributes.PublicName)
}

func TestDeleteStatusPageResource(t *testing.T) {
	err := bs.DeleteStatusPageResource(newStatusPageID, statusPageResourceID)
	assert.Nil(t, err)

	err = bs.DeleteStatusPage(newStatusPageID)
	assert.Nil(t, err)

	err = bs.DeleteMonitor(monitorID)
}

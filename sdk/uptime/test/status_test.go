package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

var statusPageID string

func TestListStatusPages(t *testing.T) {
	// Test case 1: Successful API response
	statusPages, err := bs.ListStatusPages()
	assert.Nil(t, err)
	assert.NotNil(t, statusPages)
	assert.IsType(t, &uptime.StatusPages{}, statusPages)
}

func TestCreateStatusPage(t *testing.T) {
	reqBody := uptime.StatusPageReqBody{
		CompanyName: "test_company_update",
		Subdomain:   "test-subdomain",
		Timezone:    "Casablanca",
	}
	statusPage, err := bs.CreateStatusPage(reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPage)
	assert.IsType(t, &uptime.StatusPage{}, statusPage)
	assert.Equal(t, reqBody.CompanyName, statusPage.Attributes.CompanyName)
	assert.Equal(t, reqBody.Subdomain, statusPage.Attributes.Subdomain)
	assert.Equal(t, reqBody.Timezone, statusPage.Attributes.Timezone)
	statusPageID = statusPage.ID
}

func TestGetStatusPage(t *testing.T) {
	statusPage, err := bs.GetStatusPage(statusPageID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPage)
	assert.IsType(t, &uptime.StatusPage{}, statusPage)
	assert.Equal(t, statusPageID, statusPage.ID)
}

func TestUpdateStatusPage(t *testing.T) {
	reqBody := uptime.StatusPageReqBody{
		CompanyName: "test_company_update",
		Subdomain:   "test-subdomain-update",
		Timezone:    "Montevideo",
	}
	statusPage, err := bs.UpdateStatusPage(statusPageID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPage)
	assert.IsType(t, &uptime.StatusPage{}, statusPage)
	assert.Equal(t, statusPageID, statusPage.ID)
	assert.Equal(t, reqBody.CompanyName, statusPage.Attributes.CompanyName)
	assert.Equal(t, reqBody.Subdomain, statusPage.Attributes.Subdomain)
	assert.Equal(t, reqBody.Timezone, statusPage.Attributes.Timezone)
}

func TestDeleteStatusPage(t *testing.T) {
	err := bs.DeleteStatusPage(statusPageID)
	assert.Nil(t, err)
}

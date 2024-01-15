package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
)

var (
	newStatusPageID, statusPageSectionID string
)

func TestCreateStatusPageSection(t *testing.T) {
	reqBody := models.StatusPageSectionReqBody{
		Name:     "test_section",
		Position: 0,
	}
	newStatusPage, err := bs.CreateStatusPage(models.StatusPageReqBody{
		CompanyName: "test_company_update",
		Subdomain:   "test-subdomain-section",
		Timezone:    "Casablanca",
	})
	assert.Nil(t, err)
	assert.NotNil(t, newStatusPage)
	newStatusPageID = newStatusPage.ID
	statusPageSection, err := bs.CreateStatusPageSection(newStatusPageID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageSection)
	assert.IsType(t, &models.StatusPageSection{}, statusPageSection)
	assert.Equal(t, reqBody.Name, statusPageSection.Attributes.Name)
	assert.Equal(t, reqBody.Position, statusPageSection.Attributes.Position)
	statusPageSectionID = statusPageSection.ID
}

func TestListStatusPageSections(t *testing.T) {
	// Test case 1: Successful API response
	statusPageSections, err := bs.ListStatusPageSections(newStatusPageID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageSections)
	assert.IsType(t, &models.StatusPageSections{}, statusPageSections)
}

func TestGetStatusPageSection(t *testing.T) {
	statusPageSection, err := bs.GetStatusPageSection(newStatusPageID, statusPageSectionID)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageSection)
	assert.IsType(t, &models.StatusPageSection{}, statusPageSection)
	assert.Equal(t, statusPageSectionID, statusPageSection.ID)
}

func TestUpdateStatusPageSection(t *testing.T) {
	reqBody := models.StatusPageSectionReqBody{
		Name: "test_update_section",
	}
	statusPageSection, err := bs.UpdateStatusPageSection(newStatusPageID, statusPageSectionID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, statusPageSection)
	assert.IsType(t, &models.StatusPageSection{}, statusPageSection)
	assert.Equal(t, statusPageSectionID, statusPageSection.ID)
	assert.Equal(t, reqBody.Name, statusPageSection.Attributes.Name)
}

func TestDeleteStatusPageSection(t *testing.T) {
	err := bs.DeleteStatusPageSection(newStatusPageID, statusPageSectionID)
	assert.Nil(t, err)

	err = bs.DeleteStatusPage(newStatusPageID)
	assert.Nil(t, err)
}

package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sudodeo/betterstack-go/models/logs"
)

var (
	testSourceID string
)

func TestListSources(t *testing.T) {
	page := 1
	// Test case 1: Successful API response
	sources, err := bs.ListSources(&page, nil)
	assert.Nil(t, err)
	assert.NotNil(t, sources)
	assert.IsType(t, &logs.SourcesResponse{}, sources)

	// Test case 2: failed API response due to invalid page number
	page = -1
	perPage := 50
	sources, err = bs.ListSources(&page, &perPage)
	assert.NotNil(t, err)
	assert.Nil(t, sources)
}

func TestCreateSource(t *testing.T) {
	// Test case 1: successful API response
	bodyParams := &logs.CreateSourceBodyParams{
		Name:     "test_source",
		Platform: logs.PlatformList.Docker,
	}
	source, err := bs.CreateSource(*bodyParams)
	assert.Nil(t, err)
	assert.NotNil(t, source)
	assert.IsType(t, &logs.SourceData{}, source)
	testSourceID = source.ID

	// Test case 2: failed API response due to missing name
	bodyParams.Name = ""
	source, err = bs.CreateSource(*bodyParams)
	assert.NotNil(t, err)
	assert.Nil(t, source)

	// Test case 3: failed API response due to unsupported platform
	bodyParams.Name = "unsupported_source"
	bodyParams.Platform = "invalid_platform"
	source, err = bs.CreateSource(*bodyParams)
	assert.NotNil(t, err)
	assert.Nil(t, source)

}

func TestGetSource(t *testing.T) {
	// Test case 1: Successful API response
	source, err := bs.GetSource(testSourceID)
	assert.Nil(t, err)
	assert.NotNil(t, source)
	assert.IsType(t, &logs.SourceData{}, source)

	// Test case 2: failed API response
	source, err = bs.GetSource("invalidID")
	assert.NotNil(t, err)
	assert.Nil(t, source)
}

func TestUpdateSource(t *testing.T) {
	// Test case 1: Successful API response
	updateName := "update_test_source"
	bodyParams := logs.UpdateSourceBodyParams{
		Name:            updateName,
		IngestingPaused: true,
	}
	source, err := bs.UpdateSource(testSourceID, bodyParams)
	assert.Nil(t, err)
	assert.NotNil(t, source)
	assert.IsType(t, &logs.SourceData{}, source)
	assert.Equal(t, updateName, source.Attributes.Name)
	assert.Equal(t, true, source.Attributes.IngestingPaused)

	// Test case 2: failed API response due to missing name
	bodyParams.Name = ""
	source, err = bs.UpdateSource(testSourceID, bodyParams)
	assert.NotNil(t, err)
	assert.Nil(t, source)

	// Test case 3: failed API response due to wrong source_id
	bodyParams.Name = "nonexistent_source_id"
	source, err = bs.UpdateSource("testSourceID", bodyParams)
	assert.NotNil(t, err)
	assert.Nil(t, source)
}

func TestDeleteSource(t *testing.T) {
	// Test case 1: Delete created source
	err := bs.DeleteSource(testSourceID)
	assert.Nil(t, err)

	// Test case2: Attempt to delete a non-existent source
	err = bs.DeleteSource("nonexistent_source_id")
	assert.NotNil(t, err)
}

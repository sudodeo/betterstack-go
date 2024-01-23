package uptime_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/sudodeo/betterstack-go/models"
// )

// var newRelicID string

// func TestListNewRelics(t *testing.T) {
// 	monitors, err := bs.ListNewRelics()
// 	assert.Nil(t, err)
// 	assert.NotNil(t, monitors)
// 	assert.IsType(t, &models.NewRelics{}, monitors)
// }

// func TestCreateNewRelic(t *testing.T) {
// 	reqBody := models.NewRelicReqBody{
// 		Name: "new relic name",
// 	}
// 	monitor, err := bs.CreateNewRelic(reqBody)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, monitor)
// 	assert.IsType(t, &models.NewRelic{}, monitor)
// 	assert.Equal(t, reqBody.Name, monitor.Attributes.Name)
// 	newRelicID = monitor.ID
// }

// func TestGetNewRelic(t *testing.T) {
// 	monitor, err := bs.GetNewRelic(newRelicID)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, monitor)
// 	assert.IsType(t, &models.NewRelic{}, monitor)
// 	assert.Equal(t, newRelicID, monitor.ID)
// }

// func TestUpdateNewRelic(t *testing.T) {
// 	reqBody := models.NewRelicReqBody{
// 		Name: "new relic name update",
// 	}
// 	monitor, err := bs.UpdateNewRelic(newRelicID, reqBody)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, monitor)
// 	assert.IsType(t, &models.NewRelic{}, monitor)
// 	assert.Equal(t, newRelicID, monitor.ID)
// 	assert.Equal(t, reqBody.Name, monitor.Attributes.Name)
// }

// func TestDeleteNewRelic(t *testing.T) {
// 	err := bs.DeleteNewRelic(newRelicID)
// 	assert.Nil(t, err)
// }

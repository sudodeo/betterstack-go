package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
)

var oncallCalendarID string

func TestListOncallCalendars(t *testing.T) {
	oncallCalaendars, err := bs.ListOncallCalendars()
	assert.Nil(t, err)
	assert.NotNil(t, oncallCalaendars)
	assert.IsType(t, &models.OncallCalendars{}, oncallCalaendars)
}

func TestGetOncallCalendar(t *testing.T) {
	oncallCalendarID = "192097"
	oncallCalendar, err := bs.GetOncallCalendar(oncallCalendarID)
	assert.Nil(t, err)
	assert.NotNil(t, oncallCalendar)
	assert.IsType(t, &models.OncallCalendar{}, oncallCalendar)
	assert.Equal(t, oncallCalendarID, oncallCalendar.ID)
}

package uptime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sudodeo/betterstack-go/models/uptime"
)

var incidentID string

func TestListIncidents(t *testing.T) {
	incidents, err := bs.ListIncidents()
	assert.Nil(t, err)
	assert.NotNil(t, incidents)
	assert.IsType(t, &uptime.Incidents{}, incidents)
}

func TestCreateIncident(t *testing.T) {
	reqBody := uptime.IncidentReqBody{
		RequesterEmail: "test@example.com",
		Summary:        "Test incident go",
	}
	incident, err := bs.CreateIncident(reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, incident)
	assert.IsType(t, &uptime.Incident{}, incident)
	assert.Equal(t, reqBody.Summary, incident.Attributes.Cause)
	incidentID = incident.ID
}

func TestAcknowledgeIncident(t *testing.T) {
	incident, err := bs.AcknowledgeIncident(incidentID, "test@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, incident)
	assert.IsType(t, &uptime.Incident{}, incident)
	assert.Equal(t, incidentID, incident.ID)
	incidentID = incident.ID
}

func TestResolveIncident(t *testing.T) {
	incident, err := bs.ResolveIncident(incidentID, "test@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, incident)
	assert.IsType(t, &uptime.Incident{}, incident)
	assert.Equal(t, incidentID, incident.ID)
	incidentID = incident.ID
}

func TestGetIncident(t *testing.T) {
	incident, err := bs.GetIncident(incidentID)
	assert.Nil(t, err)
	assert.NotNil(t, incident)
	assert.IsType(t, &uptime.Incident{}, incident)
	assert.Equal(t, incidentID, incident.ID)
}

func TestGetIncidentTimelineEvents(t *testing.T) {
	incidentTimelineEvents, err := bs.GetIncidentTimelineEvents(incidentID)
	assert.Nil(t, err)
	assert.NotNil(t, incidentTimelineEvents)
	assert.IsType(t, &uptime.IncidentTimelineEvents{}, incidentTimelineEvents)
}

func TestDeleteIncident(t *testing.T) {
	err := bs.DeleteIncident(incidentID)
	assert.Nil(t, err)
}

package uptime_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudodeo/betterstack-go/models"
	"testing"
)

var policyID string

func TestListPolicies(t *testing.T) {
	policies, err := bs.ListEscalationPolicies()
	assert.Nil(t, err)
	assert.NotNil(t, policies)
	assert.IsType(t, &models.EscalationPolicies{}, policies)
}

func TestCreatePolicy(t *testing.T) {
	reqBody := models.PolicyReqBody{
		Name: "TestPolicy",
	}
	policy, err := bs.CreateEscalationPolicy(reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, policy)
	assert.IsType(t, &models.EscalationPolicy{}, policy)
	assert.Equal(t, reqBody.Name, policy.Attributes.Name)
	policyID = policy.ID
}

func TestGetPolicy(t *testing.T) {
	policy, err := bs.GetEscalationPolicy(policyID)
	assert.Nil(t, err)
	assert.NotNil(t, policy)
	assert.IsType(t, &models.EscalationPolicy{}, policy)
	assert.Equal(t, policyID, policy.ID)
}

func TestUpdatePolicy(t *testing.T) {
	reqBody := models.PolicyReqBody{
		Name: "Updated_TestPolicy",
	}
	policy, err := bs.UpdateEscalationPolicy(policyID, reqBody)
	assert.Nil(t, err)
	assert.NotNil(t, policy)
	assert.IsType(t, &models.EscalationPolicy{}, policy)
	assert.Equal(t, policyID, policy.ID)
	assert.Equal(t, reqBody.Name, policy.Attributes.Name)
}

func TestDeletePolicy(t *testing.T) {
	err := bs.DeleteEscalationPolicy(policyID)
	assert.Nil(t, err)
}

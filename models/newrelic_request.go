package models

// NewRelicReqBody is the request body for creating/updating a new relic integration
type NewRelicReqBody struct {
	Name string `json:"name"` // required
	// ! *Any other params you'd like to set. See the list of all New Relic integration API parameters for details.
}

package uptime

// MetadataRecordReqBody represents the request body to create a metadata
type MetadataRecordReqBody struct {
	OwnerID   string `json:"owner_id"`        // The ID of the owner
	OwnerType string `json:"owner_type"`      // The type of the owner record - accepted values can be accessed with OwnerType
	Key       string `json:"key"`             // Metadata key
	Value     string `json:"value,omitempty"` // Optional, Metadata value
}

type owner struct {
	Monitor            string
	Heartbeat          string
	Incident           string
	WebhookIntegration string
	EmailIntegration   string
}

// OwnerType represents the type of owner for a metadata request.
var OwnerType = owner{
	Monitor:            "Monitor",
	Heartbeat:          "Heartbeat",
	Incident:           "Incident",
	WebhookIntegration: "WebhookIntegration",
	EmailIntegration:   "EmailIntegration",
}

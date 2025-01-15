package models

// Sources represents a response containing a list of log sources.
type Sources struct {
	Data       []Source   `json:"data"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// Source represents a log source.
type Source struct {
	Data
	Attributes SourcesAttributes `json:"attributes"`
}

type SourcesAttributes struct {
	TeamID                int      `json:"team_id"`
	Name                  string   `json:"name"`
	TableName             string   `json:"table_name"`
	Platform              string   `json:"platform"`
	Token                 string   `json:"token"`
	Retention             int      `json:"retention"`
	IngestingPaused       bool     `json:"ingesting_paused"`
	AutogenerateViews     bool     `json:"autogenerate_views"`
	CreatedAt             string   `json:"created_at"`
	UpdatedAt             string   `json:"updated_at"`
	ResourceGroupID       int      `json:"resource_group_id"`
	Position              int      `json:"position"`
	ReceivingLogs         bool     `json:"receiving_logs"`
	Columns               []string `json:"columns"`
	ColumnsQuota          int      `json:"columns_quota"`
	LivetailPrimaryFields []string `json:"livetail_primary_fields"`
	GrafanaPrimaryFields  []string `json:"grafana_primary_fields"`
}

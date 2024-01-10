package logs

type SourcesResponse struct {
	Data       []SourceData `json:"data"`
	Pagination Pagination   `json:"pagination,omitempty"`
}

type SourceData struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
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

type Pagination struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

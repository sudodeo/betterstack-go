package models

// CreateSourceBodyParams represents the request body required to create log source
type CreateSourceBodyParams struct {
	Name     string `json:"name"`     // source name
	Platform string `json:"platform"` // source platform (get options via `PlatformList`)
}

// UpdateSourceBodyParams represents the request body required to update log source
type UpdateSourceBodyParams struct {
	Name            string `json:"name"`             // Source name
	IngestingPaused bool   `json:"ingesting_paused"` // Enable or disable ingesting for this source
}

type Platform struct {
	Apache2           string
	AwsCloudwatch     string
	AwsEcs            string
	AwsFargate        string
	CloudflareWorker  string
	DatadogAgent      string
	Docker            string
	Dokku             string
	Dotnet            string
	ElasticSearch     string
	Filebeat          string
	Fluentbit         string
	Fluentd           string
	flyIO             string
	Haproxy           string
	Heroku            string
	HTTP              string
	Java              string
	Javascript        string
	Kubernetes        string
	Logstash          string
	Minio             string
	Mongodb           string
	MySQL             string
	Nginx             string
	OpenTelemetry     string
	PHP               string
	Postgresql        string
	Prometheus        string
	Python            string
	Rabbitmq          string
	Redis             string
	Render            string
	Rsyslog           string
	Ruby              string
	SyslogNg          string
	Traefik           string
	Ubuntu            string
	Vector            string
	VercelIntegration string
}

// PlatformList is the list of supported platforms
var PlatformList = Platform{
	Apache2:           "apache2",
	AwsCloudwatch:     "aws_cloudwatch",
	AwsEcs:            "aws_ecs",
	AwsFargate:        "aws_fargate",
	CloudflareWorker:  "cloudflare_worker",
	DatadogAgent:      "datadog_agent",
	Docker:            "docker",
	Dokku:             "dokku",
	Dotnet:            "dotnet",
	ElasticSearch:     "elasticsearch",
	Filebeat:          "filebeat",
	Fluentbit:         "fluentbit",
	Fluentd:           "fluentd",
	flyIO:             "fly_io",
	Haproxy:           "haproxy",
	Heroku:            "heroku",
	HTTP:              "http",
	Java:              "java",
	Javascript:        "javascript",
	Kubernetes:        "kubernetes",
	Logstash:          "logstash",
	Minio:             "minio",
	Mongodb:           "mongodb",
	MySQL:             "mysql",
	Nginx:             "nginx",
	OpenTelemetry:     "open_telemetry",
	PHP:               "php",
	Postgresql:        "postgresql",
	Prometheus:        "prometheus",
	Python:            "python",
	Rabbitmq:          "rabbitmq",
	Redis:             "redis",
	Render:            "render",
	Rsyslog:           "rsyslog",
	Ruby:              "ruby",
	SyslogNg:          "syslog-ng",
	Traefik:           "traefik",
	Ubuntu:            "ubuntu",
	Vector:            "vector",
	VercelIntegration: "vercel_integration",
}

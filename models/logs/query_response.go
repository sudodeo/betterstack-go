package logs

// Logs represents a response containing a list of logs
type Logs struct {
	Data       []log      `json:"data"`
	Pagination pagination `json:"pagination"`
}

type log struct {
	App                          string `json:"_app"`
	SourceID                     string `json:"_source_id"`
	Dt                           string `json:"dt"`
	DtUnix                       string `json:"_dt"`
	EventHTTPRequestReceivedPort any    `json:"event.http_request_received.port"`
	EventHTTPResponseSentStatus  any    `json:"event.http_response_sent.status"`
	Message                      string `json:"message"`
	Level                        string `json:"level"`
	ContextRuntimeThreadID       string `json:"context.runtime.thread_id"`
	ContextSystemPid             string `json:"context.system.pid"`
	ContextHTTPPath              any    `json:"context.http.path"`
	ContextHTTPMethod            any    `json:"context.http.method"`
	ContextHTTPHost              any    `json:"context.http.host"`
	InsertIndex                  int    `json:"_insert_index"`
	JSON                         string `json:"json"`
}

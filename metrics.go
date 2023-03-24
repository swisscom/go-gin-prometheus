package promgin

var requestCount = &Metric{
	Name:        "requests_total",
	Description: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	Type:        "counter_vec",
	Args:        []string{"code", "method", "handler", "is_aborted", "is_websocket", "content_type", "proto", "host", "url"}}

var requestDurationMs = &Metric{
	Name:        "request_duration_seconds",
	Description: "The HTTP request latencies in seconds",
	Type:        "histogram_vec",
	Args:        []string{"code", "method", "url"},
}

var responseSizeBytes = &Metric{
	Name:        "response_size_bytes",
	Description: "The HTTP response sizes in bytes.",
	Type:        "summary"}

var requestSizeBytes = &Metric{
	Name:        "request_size_bytes",
	Description: "The HTTP request sizes in bytes.",
	Type:        "summary"}

var standardMetrics = []*Metric{
	requestCount,
	requestDurationMs,
	responseSizeBytes,
	requestSizeBytes,
}

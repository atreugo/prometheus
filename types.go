package prometheus

// Config is the configuration for the Prometheus plugin.
type Config struct {
	// URL is the path to the metrics endpoint.
	//
	// Default: "/metrics"
	URL string

	// Method is the HTTP method to use for the metrics endpoint.
	//
	// Default: "GET"
	Method string
}

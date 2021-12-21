package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/savsgio/atreugo/v11"
)

// Register registers the prometheus plugin to the server with the given configuration.
func Register(s *atreugo.Atreugo, cfg Config) {
	if cfg.Method == "" {
		cfg.Method = defaultMethod
	}

	if cfg.URL == "" {
		cfg.URL = defaultURL
	}

	s.NetHTTPPath(cfg.Method, cfg.URL, promhttp.Handler())
}

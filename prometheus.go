package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/savsgio/atreugo/v11"
)

// Register registers the prometheus plugin with the given configuration and server.
func Register(cfg Config, s *atreugo.Atreugo) {
	if cfg.URL == "" {
		cfg.URL = defaultURL
	}

	if cfg.Method == "" {
		cfg.Method = defaultMethod
	}

	s.NetHTTPPath("GET", cfg.URL, promhttp.Handler())
}

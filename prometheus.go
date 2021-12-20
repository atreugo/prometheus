package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/savsgio/atreugo/v11"
)

// Register registers the prometheus plugin with the given configuration and server.
func Register(cfg Config, s *atreugo.Atreugo) {
	if cfg.Method == "" {
		cfg.Method = defaultMethod
	}

	if cfg.URL == "" {
		cfg.URL = defaultURL
	}

	s.NetHTTPPath(cfg.Method, cfg.URL, promhttp.Handler())
}

package prometheus

import "github.com/valyala/fasthttp"

const (
	defaultURL    = "/metrics"
	defaultMethod = fasthttp.MethodGet
)

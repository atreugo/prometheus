package prometheus

import (
	"testing"

	"github.com/atreugo/httptest/assert"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func Test_Register(t *testing.T) {
	s := atreugo.New(atreugo.Config{})
	Register(s, Config{})

	req := new(fasthttp.Request)
	req.SetRequestURI(defaultURL)
	req.Header.SetMethod(defaultMethod)

	assert.Server(t, req, s, func(resp *fasthttp.Response) {
		if statusCode := resp.StatusCode(); statusCode != fasthttp.StatusOK {
			t.Errorf("status code %d, want %d", statusCode, fasthttp.StatusOK)
		}

		if len(resp.Body()) == 0 {
			t.Error("body is empty")
		}
	})
}

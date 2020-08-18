package middleware

import (
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

// TracingMiddleware is a wrapper around github.com/opentracing-contrib/go-stdlib/nethttp
type TracingMiddleware struct {
	tracer  opentracing.Tracer
	options []nethttp.MWOption
}

// NewPrometheus returns a new prometheus MetricsMiddleware handler.
func NewTracing(tracer opentracing.Tracer, options ...nethttp.MWOption) func(next http.Handler) http.Handler {
	var t TracingMiddleware
	t.tracer = tracer
	t.options = options
	return t.handler
}

func (t TracingMiddleware) handler(next http.Handler) http.Handler {
	return nethttp.Middleware(t.tracer, next, t.options...)
}

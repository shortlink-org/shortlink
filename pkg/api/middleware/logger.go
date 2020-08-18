package middleware

import (
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/batazor/shortlink/internal/logger"

	"github.com/go-chi/chi/middleware"
)

type chilogger struct { // nolint unused
	logZ logger.Logger
}

// Logger returns a new Zap Middleware handler.
func Logger(log logger.Logger) func(next http.Handler) http.Handler { // nolint unused
	return chilogger{
		log,
	}.middleware
}

func (c chilogger) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		// Get span ID
		span := opentracing.SpanFromContext(r.Context())
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()

		latency := time.Since(start)

		var fields = logger.Fields{
			"status":  ww.Status(),
			"took":    latency,
			"remote":  r.RemoteAddr,
			"request": r.RequestURI,
			"method":  r.Method,
		}
		if traceID != "" {
			fields["traceID"] = traceID
		}
		c.logZ.Info("request completed", fields)
	}
	return http.HandlerFunc(fn)
}

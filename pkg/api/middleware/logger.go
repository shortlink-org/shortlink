package middleware

import (
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"

	"github.com/go-chi/chi/v5/middleware"
)

type chilogger struct { // nolint unused
	logZ logger.Logger
}

// Logger returns a new Zap Middleware handler.
func Logger(log logger.Logger) func(next http.Handler) http.Handler {
	return chilogger{
		log,
	}.middleware
}

func (c chilogger) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		latency := time.Since(start)

		var fields = field.Fields{
			"status":  ww.Status(),
			"took":    latency,
			"remote":  r.RemoteAddr,
			"request": r.RequestURI,
			"method":  r.Method,
		}

		// Get span ID
		span := opentracing.SpanFromContext(r.Context())
		if span != nil {
			traceID := span.Context().(jaeger.SpanContext).TraceID().String()
			fields["traceID"] = traceID
		}

		c.logZ.Info("request completed", fields)
	}
	return http.HandlerFunc(fn)
}

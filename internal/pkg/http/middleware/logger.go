package middleware

import (
	"net/http"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"

	"github.com/go-chi/chi/v5/middleware"
)

type chilogger struct {
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

		fields := field.Fields{
			"status":  ww.Status(),
			"took":    latency,
			"remote":  r.RemoteAddr,
			"request": r.RequestURI,
			"method":  r.Method,
		}

		// Get span ID
		span := trace.LinkFromContext(r.Context()).SpanContext
		fields["traceID"] = span.TraceID().String()

		c.logZ.Info("request completed", fields)
	}

	return http.HandlerFunc(fn)
}

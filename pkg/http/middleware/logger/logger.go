package logger_middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

type chilogger struct {
	log logger.Logger
}

// Logger returns a new Zap Middleware handler.
func Logger(log logger.Logger) func(next http.Handler) http.Handler {
	return chilogger{
		log: log,
	}.middleware
}

func (c chilogger) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			latency := time.Since(start)

			fields := field.Fields{
				"status":  ww.Status(),
				"took":    latency,
				"remote":  r.RemoteAddr,
				"request": r.RequestURI,
				"method":  r.Method,
			}

			c.log.InfoWithContext(r.Context(), "request completed", fields)
		}()

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}

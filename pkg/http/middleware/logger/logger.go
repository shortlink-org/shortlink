package logger_middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/shortlink-org/go-sdk/logger"
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

			c.log.InfoWithContext(r.Context(), "request completed",
				slog.Int("status", ww.Status()),
				slog.Duration("took", latency),
				slog.String("remote", r.RemoteAddr),
				slog.String("request", r.RequestURI),
				slog.String("method", r.Method),
			)
		}()

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}

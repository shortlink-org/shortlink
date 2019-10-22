package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type chilogger struct {
	logZ *zap.Logger
}

// NewZapMiddleware returns a new Zap Middleware handler.
func Logger(logger *zap.Logger) func(next http.Handler) http.Handler {
	return chilogger{
		logZ: logger,
	}.middleware
}

func (c chilogger) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var requestID string
		if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
			requestID = reqID.(string)
		}
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		latency := time.Since(start)

		fields := []zapcore.Field{
			zap.Int("status", ww.Status()),
			zap.Duration("took", latency),
			zap.String("remote", r.RemoteAddr),
			zap.String("request", r.RequestURI),
			zap.String("method", r.Method),
		}
		if requestID != "" {
			fields = append(fields, zap.String("request-id", requestID))
		}
		c.logZ.Info("request completed", fields...)
	}
	return http.HandlerFunc(fn)
}

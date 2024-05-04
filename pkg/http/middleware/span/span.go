package span_middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel/trace"
)

const (
	// TraceIDHeader is the header key for the trace id.
	TraceIDHeader = "trace_id"
)

type span struct{}

// Span is a middleware that adds a span to the response context.
func Span() func(next http.Handler) http.Handler {
	return span{}.middleware
}

func (s span) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// Check if "trace_id" already exists in the header
		if ww.Header().Get(TraceIDHeader) == "" {
			// Inject spanId in response header
			ww.Header().Add(TraceIDHeader, trace.SpanFromContext(r.Context()).SpanContext().TraceID().String())
		}

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}

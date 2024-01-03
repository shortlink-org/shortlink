package span_middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel/trace"
)

type span struct{}

// Span is a middleware that adds a span to the response context.
func Span() func(next http.Handler) http.Handler {
	return span{}.middleware
}

func (s span) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// Inject spanId in response header
		ww.Header().Add("trace_id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())
		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}

package span_middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpanMiddleware(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "span")
	t.Attr("component", "http")

		t.Attr("type", "unit")
		t.Attr("package", "span")
		t.Attr("component", "http")
	
	// Test handler that does nothing
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Apply the Span middleware
	middleware := Span()(testHandler)

	// Create an HTTP request
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)

	// Using a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using our middleware
	middleware.ServeHTTP(rr, req)

	// Assert the trace_id is in the response header
	traceID := rr.Header().Get(TraceIDHeader)
	require.NotEmpty(t, traceID)
}

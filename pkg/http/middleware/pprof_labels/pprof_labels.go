package pprof_labels_middleware

import (
	"net/http"
	"runtime/pprof"
)

// Labels is a middleware function that adds pprof labels to the context of the incoming HTTP request.
// These labels include the request path and method.
// The updated context is then used to serve the HTTP request.
func Labels(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := pprof.WithLabels(r.Context(), pprof.Labels(
			"path", r.URL.Path,
			"method", r.Method,
		))

		pprof.SetGoroutineLabels(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

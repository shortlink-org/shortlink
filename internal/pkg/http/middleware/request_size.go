package middleware

import (
	"net/http"
)

// RequestSize is a middleware that will limit request sizes to a specified
// number of bytes. It uses MaxBytesReader to do so.
//
// Size:
// 1<<10 - 1KB
// 1<<20 - 1MB
// 1<<30 - 1GB
func RequestSize(bytes int64) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, bytes)
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}

	return f
}

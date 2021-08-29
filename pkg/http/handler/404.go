package handler

import (
	"net/http"
)

// NotFoundHandler - default handler for don't existing routers
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{}`))
}

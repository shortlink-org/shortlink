package order

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	jsonpb protojson.MarshalOptions
)

// Routes creates a REST router
func Routes(r chi.Router) {
	r.Get("/order/{hash}", Get)
	r.Get("/orders", List)
	r.Post("/order", Add)
	r.Delete("/order/{hash}", Delete)
}

// Add ...
func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// Get ...
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// List ...
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

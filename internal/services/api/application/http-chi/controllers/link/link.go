package link_api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/batazor/shortlink/internal/services/api/application/http-chi/helpers"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
)

var (
	jsonpb protojson.MarshalOptions
)

type Handler struct {
	LinkServiceClient link_rpc.LinkServiceClient
}

// Routes creates a REST router
func Routes(
	link_rpc link_rpc.LinkServiceClient,
) chi.Router {
	r := chi.NewRouter()

	h := &Handler{
		LinkServiceClient: link_rpc,
	}

	r.Get("/{hash}", h.Get)
	r.Get("/list", h.List)
	r.Post("/", h.Add)
	r.Put("/", h.Update)
	r.Patch("/", h.Update)
	r.Delete("/{hash}", h.Delete)

	return r
}

// Add ...
func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var request v1.Link
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	// Save link
	response, err := h.LinkServiceClient.Add(r.Context(), &link_rpc.AddRequest{Link: &request})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint errcheck
}

// Update ...
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var request v1.Link
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	// Update link
	response, err := h.LinkServiceClient.Update(r.Context(), &link_rpc.UpdateRequest{Link: &request})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint errcheck
}

// Get ...
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	response, err := h.LinkServiceClient.Get(r.Context(), &link_rpc.GetRequest{Hash: hash})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	var errorLink *v1.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := jsonpb.Marshal(response.GetLink())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// List ...
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Get filter
	filter := r.URL.Query().Get("filter")

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	response, err := h.LinkServiceClient.List(r.Context(), &link_rpc.ListRequest{Filter: filter})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	var errorLink *v1.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := jsonpb.Marshal(response.GetLinks())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// Delete ...
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	_, err := h.LinkServiceClient.Delete(r.Context(), &link_rpc.DeleteRequest{Hash: hash})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`))
}

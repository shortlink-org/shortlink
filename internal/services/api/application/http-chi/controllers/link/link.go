package link_api

import (
	"errors"
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
)

var jsonpb protojson.MarshalOptions

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
	r.Get("/", h.List)
	r.Post("/", h.Add)
	r.Put("/", h.Update)
	r.Patch("/", h.Update)
	r.Delete("/{hash}", h.Delete)

	return r
}

// Add ...
// @Summary Add link
// @Description Add link
// @ID add-link
// @Accept  json
// @Produce  json
// @Group Links
// @Success 200 {object} link_rpc.AddResponse
// @Router /links [post]
// @Param link body link_rpc.AddRequest true "Link"
func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var request v1.Link
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	// Save link
	response, err := h.LinkServiceClient.Add(r.Context(), &link_rpc.AddRequest{Link: &request})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint:errcheck
}

// Update ...
// @Summary Update link
// @Description Update link
// @ID update-link
// @Accept  json
// @Produce  json
// @Group Links
// @Success 200 {object} link_rpc.UpdateResponse
// @Router /links/:hash [put]
// @Param link body link_rpc.UpdateRequest true "Link"
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var request v1.Link
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	// Update link
	response, err := h.LinkServiceClient.Update(r.Context(), &link_rpc.UpdateRequest{Link: &request})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint:errcheck
}

// Get ...
// @Summary Get link
// @Description Get link
// @ID get-link
// @Accept  json
// @Produce  json
// @Group Links
// @Success 200 {object} link_rpc.GetResponse
// @NotFound 404 {object} link_rpc.GetResponse
// @Router /links/{hash} [get]
// @Param hash path string true "Hash"
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	hash := chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint:errcheck

		return
	}

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	response, err := h.LinkServiceClient.Get(r.Context(), &link_rpc.GetRequest{Hash: hash})
	if err != nil {
		var errorLink *v1.NotFoundError

		if errors.Is(err, errorLink) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

			return
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.GetLink())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint:errcheck
}

// List ...
// @Summary List links
// @Description List links
// @ID list-links
// @Accept  json
// @Produce  json
// @Group Links
// @Success 200 {object} link_rpc.ListResponse
// @Router /links [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Get filter
	filter := r.URL.Query().Get("filter")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	response, err := h.LinkServiceClient.List(r.Context(), &link_rpc.ListRequest{Filter: filter})
	if err != nil {
		var errorLink *v1.NotFoundError

		if errors.Is(err, errorLink) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

			return
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.GetLinks())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint:errcheck
}

// Delete ...
// @Summary Delete link
// @Description Delete link
// @ID delete-link
// @Accept  json
// @Produce  json
// @Group Links
// @Success 200 ""
// @Router /links/{hash} [delete]
// @Param hash path string true "Hash"
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	hash := chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint:errcheck

		return
	}

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	_, err := h.LinkServiceClient.Delete(r.Context(), &link_rpc.DeleteRequest{Hash: hash})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint:errcheck
}

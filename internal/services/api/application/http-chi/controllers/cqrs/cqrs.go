package cqrs_api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/batazor/shortlink/internal/services/api/application/http-chi/helpers"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	link_cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
)

var (
	jsonpb protojson.MarshalOptions
)

type Handler struct {
	LinkCommandServiceClient link_cqrs.LinkCommandServiceClient
	LinkQueryServiceClient   link_cqrs.LinkQueryServiceClient
}

// Routes creates a REST router
func Routes(
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
) chi.Router {
	r := chi.NewRouter()

	h := &Handler{
		LinkCommandServiceClient: link_command,
		LinkQueryServiceClient:   link_query,
	}

	r.Get("/link/{hash}", h.GetByCQRS)

	return r
}

// GetByCQRS ...
func (h *Handler) GetByCQRS(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	response, err := h.LinkQueryServiceClient.Get(r.Context(), &link_cqrs.GetRequest{Hash: hash})
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

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

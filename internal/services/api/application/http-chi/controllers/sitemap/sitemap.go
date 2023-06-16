package sitemap_api

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/trace"

	v1 "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
)

type Handler struct {
	SitemapServiceClient v1.SitemapServiceClient
}

// Routes creates a REST router
func Routes(sitemap_rpc v1.SitemapServiceClient) chi.Router {
	r := chi.NewRouter()

	h := &Handler{
		SitemapServiceClient: sitemap_rpc,
	}

	r.Post("/", h.Parse)

	return r
}

// Parse ...
func (h *Handler) Parse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var request v1.ParseRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	// Parse link
	_, err = h.SitemapServiceClient.Parse(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{}`)) // nolint:errcheck
}

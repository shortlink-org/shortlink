package sitemap

import (
	"net/http"

	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"
	sitemapv1 "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/infrastructure/rpc/sitemap/v1"
	"github.com/segmentio/encoding/json"

	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/api"
)

type SitemapController struct {
	SitemapServiceClient sitemap_rpc.SitemapServiceClient
}

// AddSitemap triggers sitemap parsing.
func (c *SitemapController) AddSitemap(w http.ResponseWriter, r *http.Request) {
	// Parse request
	var request api.AddSitemapRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "invalid request payload"}`)) //nolint:errcheck

		return
	}

	if request.Url == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "url is required"}`)) //nolint:errcheck
		return
	}

	if c.SitemapServiceClient == nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "sitemap service unavailable"}`)) //nolint:errcheck
		return
	}

	_, err = c.SitemapServiceClient.Parse(r.Context(), &sitemapv1.ParseRequest{
		Url: request.Url,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "failed to parse sitemap"}`)) //nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
}

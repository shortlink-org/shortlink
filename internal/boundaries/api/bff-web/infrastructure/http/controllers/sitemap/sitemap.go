package sitemap

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/shortlink-org/shortlink/internal/boundaries/api/bff-web/infrastructure/http/api"
	v1 "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/sitemap/v1"
)

type SitemapController struct {
	SitemapServiceClient v1.SitemapServiceClient
}

// Parse - parse sitemap
func (c *SitemapController) AddSitemap(w http.ResponseWriter, r *http.Request) {
	// Parse request
	var request api.AddSitemapRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	// Parse link
	_, err = c.SitemapServiceClient.Parse(r.Context(), &v1.ParseRequest{
		Url: request.Url,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
}

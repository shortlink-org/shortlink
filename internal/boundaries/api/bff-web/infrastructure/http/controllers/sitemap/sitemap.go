package sitemap

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	v1 "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/sitemap/v1"
)

type SitemapController struct {
	SitemapServiceClient v1.SitemapServiceClient
}

// Parse - parse sitemap
func (c *SitemapController) Parse(w http.ResponseWriter, r *http.Request, params any) {
	// Parse request
	var request v1.ParseRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	// Parse link
	_, err = c.SitemapServiceClient.Parse(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusCreated)
}

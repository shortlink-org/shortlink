package metrics_middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
)

func Test_NewMetrics(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "metrics")
	t.Attr("component", "http")

		t.Attr("type", "unit")
		t.Attr("package", "metrics")
		t.Attr("component", "http")
	
	// middlewares
	middlewares, err := NewMetrics()
	require.NoError(t, err)

	// Create a new HTTP router with the Prometheus middleware
	router := chi.NewRouter()
	router.Use(middlewares)

	// Create a test request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/users/bob", http.NoBody)
	w := httptest.NewRecorder()

	// Add a test endpoint that returns a 200 OK status code
	router.Get("/users/{firstName}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Send the test request to the router
	router.ServeHTTP(w, req)

	// get the metrics from the registry
	resp, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	// Iterate over the metrics and check their values
	for _, mf := range resp {
		switch mf.GetName() {
		case "http_requests_total":
			for _, m := range mf.GetMetric() {
				labels := m.GetLabel()
				require.Equal(t, "200", getValueForLabel(labels, "status"))
				require.Equal(t, "GET", getValueForLabel(labels, "method"))
				require.Equal(t, "/users/{firstName}", getValueForLabel(labels, "path"))
				require.Equal(t, float64(1), m.GetCounter().GetValue())
			}
		case "http_request_duration_milliseconds":
			for _, m := range mf.GetMetric() {
				labels := m.GetLabel()
				require.Equal(t, "200", getValueForLabel(labels, "status"))
				require.Equal(t, "GET", getValueForLabel(labels, "method"))
				require.Equal(t, "/users/{firstName}", getValueForLabel(labels, "path"))
				require.Equal(t, uint64(1), m.GetHistogram().GetBucket()[0].GetCumulativeCount())
				require.Equal(t, uint64(1), m.GetHistogram().GetBucket()[1].GetCumulativeCount())
				require.Equal(t, uint64(1), m.GetHistogram().GetBucket()[2].GetCumulativeCount())
			}
		}
	}
}

func getValueForLabel(labels []*io_prometheus_client.LabelPair, labelName string) string {
	for _, l := range labels {
		if l.GetName() == labelName {
			return l.GetValue()
		}
	}

	return ""
}

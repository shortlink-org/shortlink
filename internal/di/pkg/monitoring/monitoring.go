package monitoring

import (
	"net/http"

	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

// New - Monitoring endpoints
func New(log logger.Logger) (*http.ServeMux, error) {
	// Create a new Prometheus registry
	registry := prometheus.NewRegistry()

	// Add Go module build info.
	err := prometheus.Register(collectors.NewBuildInfoCollector())
	if err != nil {
		return nil, err
	}

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(registry, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100)) // nolint:gomnd

	// Create a "common" listener
	commonMux := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	commonMux.Handle("/metrics", promhttp.Handler())

	// Expose a liveness check on /live
	commonMux.HandleFunc("/live", health.LiveEndpoint)

	// Expose a readiness check on /ready
	commonMux.HandleFunc("/ready", health.ReadyEndpoint)

	go func() {
		err := http.ListenAndServe("0.0.0.0:9090", commonMux)
		if err != nil {
			log.Error(err.Error())
		}
	}()
	log.Info("Run monitoring", field.Fields{
		"addr": "0.0.0.0:9090",
	})

	return commonMux, nil
}

package monitoring

import (
	"context"
	"net/http"

	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv/v1.21.0"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

type Monitoring struct {
	Handler  *http.ServeMux
	Registry *prometheus.Registry
	Metrics  *metric.MeterProvider
}

// New - Monitoring endpoints
func New(ctx context.Context, log logger.Logger) (*Monitoring, func(), error) {
	monitoring := &Monitoring{}

	// This reader is used as a stand-in for a reader that will actually export
	// data. See exporters in the go.opentelemetry.io/otel/exporters package
	// for more information.
	reader := metric.NewManualReader()

	// See the go.opentelemetry.io/otel/sdk/resource package for more
	// information about how to create and use Resources.
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(viper.GetString("SERVICE_NAME")),
		semconv.ServiceVersion(viper.GetString("SERVICE_VERSION")),
	)

	monitoring.Metrics = metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(reader),
	)

	otel.SetMeterProvider(monitoring.Metrics)

	// Create a new Prometheus registry
	monitoring.Registry = prometheus.NewRegistry()

	// Add Go module build info.
	err := prometheus.Register(collectors.NewBuildInfoCollector())
	if err != nil {
		return nil, nil, err
	}

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(monitoring.Registry, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100)) // nolint:gomnd

	// Create a "common" listener
	monitoring.Handler = http.NewServeMux()

	// Expose prometheus metrics on /metrics
	monitoring.Handler.Handle("/metrics", promhttp.HandlerFor(
		monitoring.Registry,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics, e.g., to support exemplars.
			EnableOpenMetrics: true,
		},
	))

	// Expose a liveness check on /live
	monitoring.Handler.HandleFunc("/live", health.LiveEndpoint)

	// Expose a readiness check on /ready
	monitoring.Handler.HandleFunc("/ready", health.ReadyEndpoint)

	go func() {
		err := http.ListenAndServe("0.0.0.0:9090", monitoring.Handler)
		if err != nil {
			log.Error(err.Error())
		}
	}()
	log.Info("Run monitoring", field.Fields{
		"addr": "0.0.0.0:9090",
	})

	return monitoring, func() {
		errShutdown := monitoring.Metrics.Shutdown(ctx)
		if errShutdown != nil {
			log.ErrorWithContext(ctx, errShutdown.Error())
		}
	}, nil
}

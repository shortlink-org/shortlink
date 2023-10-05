package monitoring

import (
	"context"
	"net/http"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/common"
)

type Monitoring struct {
	Handler  *http.ServeMux
	Registry *prometheus.Registry
	Metrics  *metric.MeterProvider
}

// New - Monitoring endpoints
func New(ctx context.Context, log logger.Logger) (*Monitoring, func(), error) {
	var err error
	monitoring := &Monitoring{}

	// Create a "common" meter provider for metrics
	monitoring.Metrics, err = SetMetrics()

	// Create a new Prometheus registry
	monitoring.Registry, err = SetPrometheus()
	if err != nil {
		return nil, nil, err
	}

	// Create a "common" listener
	monitoring.Handler, err = SetHandler(monitoring.Registry)
	if err != nil {
		return nil, nil, err
	}

	// Register the runtime metrics collection.
	err = SetRuntime(monitoring.Metrics)

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

// SetPrometheus - Create a new Prometheus registry
func SetPrometheus() (*prometheus.Registry, error) {
	registry := prometheus.NewRegistry()

	// Add Go module build info.
	err := prometheus.Register(collectors.NewBuildInfoCollector())
	if err != nil {
		return nil, err
	}

	return registry, nil
}

// SetMetrics - Create a "common" meter provider for metrics
func SetMetrics() (*metric.MeterProvider, error) {
	// This reader is used as a stand-in for a reader that will actually export
	// data. See exporters in the go.opentelemetry.io/otel/exporters package
	// for more information.
	reader := metric.NewManualReader()

	// See the go.opentelemetry.io/otel/sdk/resource package for more
	// information about how to create and use Resources.
	// Setup resource.
	res, err := common.NewResource(viper.GetString("SERVICE_NAME"), viper.GetString("SERVICE_VERSION"))
	if err != nil {
		return nil, err
	}

	metrics := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(reader),
	)

	otel.SetMeterProvider(metrics)

	return metrics, nil
}

// SetHandler - Create a "common" handler for metrics
func SetHandler(registry *prometheus.Registry) (*http.ServeMux, error) {
	// Create a "common" listener
	handler := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	handler.Handle("/metrics", promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics, e.g., to support exemplars.
			EnableOpenMetrics: true,
		},
	))

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(registry, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100)) // nolint:gomnd

	// Expose a liveness check on /live
	handler.HandleFunc("/live", health.LiveEndpoint)

	// Expose a readiness check on /ready
	handler.HandleFunc("/ready", health.ReadyEndpoint)

	return handler, nil
}

// SetRuntime - Register the runtime metrics collection.
func SetRuntime(metric *metric.MeterProvider) error {
	options := []runtime.Option{
		runtime.WithMinimumReadMemStatsInterval(time.Second * 1),
	}

	if metric.MeterProvider != nil {
		options = append(options, runtime.WithMeterProvider(metric))
	}

	err := runtime.Start(options...)
	if err != nil {
		return err
	}

	return nil
}

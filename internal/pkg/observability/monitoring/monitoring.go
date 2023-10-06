package monitoring

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/sdk/metric"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/common"
)

type Monitoring struct {
	Handler  *http.ServeMux
	Registry *prometheus.Exporter
	Metrics  *api.MeterProvider
}

// New - Monitoring endpoints
func New(ctx context.Context, log logger.Logger) (*Monitoring, func(), error) {
	var err error
	monitoring := &Monitoring{}

	// Create a "common" meter provider for metrics
	monitoring.Metrics, err = SetMetrics(ctx)

	// Create a "common" listener
	monitoring.Handler, err = SetHandler()
	if err != nil {
		return nil, nil, err
	}

	go func() {
		err := http.ListenAndServe("0.0.0.0:9090", monitoring.Handler)
		if err != nil {
			log.Error(err.Error())
		}
	}()
	log.Info("Run monitoring", field.Fields{
		"addr": "0.0.0.0:9090",
	})

	// Create a new OTLP exporter for sending metrics to the OpenTelemetry Collector.
	_, err = otlpmetricgrpc.New(ctx)
	if err != nil {
		return nil, nil, err
	}

	return monitoring, func() {
		errShutdown := monitoring.Metrics.Shutdown(ctx)
		if errShutdown != nil {
			log.ErrorWithContext(ctx, errShutdown.Error())
		}
	}, nil
}

// SetMetrics - Create a "common" meter provider for metrics
func SetMetrics(ctx context.Context) (*api.MeterProvider, error) {
	// See the go.opentelemetry.io/otel/sdk/resource package for more
	// information about how to create and use Resources.
	// Setup resource.
	res, err := common.NewResource(ctx, viper.GetString("SERVICE_NAME"), viper.GetString("SERVICE_VERSION"))
	if err != nil {
		return nil, err
	}

	// prometheus.DefaultRegisterer is used by default
	// so that metrics are available via promhttp.Handler.
	registry, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	provider := api.NewMeterProvider(
		api.WithResource(res),
		api.WithReader(registry),
	)

	otel.SetMeterProvider(provider)

	return provider, nil
}

// SetHandler - Create a "common" handler for metrics
func SetHandler() (*http.ServeMux, error) {
	// Create a "common" listener
	handler := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	handler.Handle("/metrics", promhttp.Handler())

	// Expose a liveness check on /live
	// TODO: recovery
	handler.Handle("/live", promhttp.Handler())

	// Expose a readiness check on /ready
	// TODO: recovery
	handler.Handle("/ready", promhttp.Handler())

	return handler, nil
}

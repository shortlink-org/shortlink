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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	promExporter "go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"

	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/observability/common"
)

type Monitoring struct {
	Handler    *http.ServeMux
	Prometheus *prometheus.Registry
	Metrics    *api.MeterProvider
}

// New - Monitoring endpoints
func New(ctx context.Context, log logger.Logger, tracer trace.TracerProvider) (*Monitoring, func(), error) {
	var err error
	monitoring := &Monitoring{}

	// Create a "common" meter provider for metrics
	monitoring.Metrics, err = monitoring.SetMetrics(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Create a "common" listener
	monitoring.Handler, err = monitoring.SetHandler()
	if err != nil {
		return nil, nil, err
	}

	go func() {
		// Create a new HTTP server for Prometheus metrics
		config := http_server.Config{
			Port:    9090,             //nolint:mnd // port for Prometheus metrics
			Timeout: 30 * time.Second, //nolint:mnd // timeout for Prometheus metrics
		}
		server := http_server.New(ctx, monitoring.Handler, config, tracer)

		errListenAndServe := server.ListenAndServe()
		if errListenAndServe != nil {
			log.Error(errListenAndServe.Error())
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
func (m *Monitoring) SetMetrics(ctx context.Context) (*api.MeterProvider, error) {
	// See the go.opentelemetry.io/otel/sdk/resource package for more
	// information about how to create and use Resources.
	// Setup resource.
	res, err := common.NewResource(ctx, viper.GetString("SERVICE_NAME"), viper.GetString("SERVICE_VERSION"))
	if err != nil {
		return nil, err
	}

	// Create a new Prometheus registry
	err = m.SetPrometheus()
	if err != nil {
		return nil, err
	}

	// prometheus.DefaultRegisterer is used by default
	// so that metrics are available via promhttp.Handler.
	registry, err := promExporter.New(
		promExporter.WithRegisterer(m.Prometheus),
	)
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
func (m *Monitoring) SetHandler() (*http.ServeMux, error) {
	// Create a "common" listener
	handler := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	handler.Handle("/metrics", promhttp.HandlerFor(
		m.Prometheus,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(m.Prometheus, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	// TODO: research problem with prometheus
	// health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100)) //nolint:mnd

	// Expose a liveness check on /live
	handler.HandleFunc("/live", health.LiveEndpoint)

	// Expose a readiness check on /ready
	handler.HandleFunc("/ready", health.ReadyEndpoint)

	return handler, nil
}

// SetPrometheus - Create a new Prometheus registry
func (m *Monitoring) SetPrometheus() error {
	m.Prometheus = prometheus.NewRegistry()

	// Add Go module build info.
	err := prometheus.Register(collectors.NewBuildInfoCollector())
	if err != nil {
		return err
	}

	return nil
}

package temporal

import (
	"github.com/uber-go/tally/v4/prometheus"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/opentelemetry"
	sdktally "go.temporal.io/sdk/contrib/tally"
	"go.temporal.io/sdk/interceptor"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
)

func New(log logger.Logger, monitor *monitoring.Monitoring) (client.Client, error) {
	metrics, err := newPrometheusScope(&prometheus.Configuration{
		ListenAddress: "0.0.0.0:9090",
		TimerType:     "histogram",
	}, monitor, log)
	if err != nil {
		return nil, err
	}

	// create Interceptor
	tracingInterceptor, err := opentelemetry.NewTracingInterceptor(opentelemetry.TracerOptions{})
	if err != nil {
		return nil, err
	}

	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		MetricsHandler: sdktally.NewMetricsHandler(metrics),
		Interceptors:   []interceptor.ClientInterceptor{tracingInterceptor},
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

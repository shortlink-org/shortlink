package auth

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

func New(log logger.Logger, tracer trace.TracerProvider, monitor *metrics.Monitoring) (*authzed.Client, error) {
	viper.SetDefault("SPICE_DB_COMMON_KEY", "secret-shortlink-preshared-key")
	viper.SetDefault("SPICE_DB_TIMEOUT", "5s")

	config, err := rpc.SetClientConfig(tracer, monitor, log)
	if err != nil {
		return nil, &ConfigurationError{Cause: err}
	}

	options := config.GetOptions()
	options = append(options,
		grpc.WithPerRPCCredentials(insecureMetadataCreds{"authorization": "Bearer " + viper.GetString("SPICE_DB_COMMON_KEY")}),
		grpc.WithIdleTimeout(viper.GetDuration("SPICE_DB_TIMEOUT")))

	client, err := authzed.NewClient(config.GetURI(), options...)
	if err != nil {
		return nil, &ClientInitError{Cause: err}
	}

	return client, nil
}

package auth

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/rpc"
)

func New(log logger.Logger, tracer trace.TracerProvider, monitoring *monitoring.Monitoring) (*authzed.Client, error) {
	viper.SetDefault("SPICE_DB_API", "shortlink.spicedb-operator:50051")
	viper.SetDefault("SPICE_DB_COMMON_KEY", "secret-shortlink-preshared-key")
	viper.SetDefault("SPICE_DB_TIMEOUT", "5s")

	config, err := rpc.SetClientConfig(tracer, monitoring, log)
	if err != nil {
		return nil, err
	}

	options := config.GetOptions()
	options = append(options,
		grpc.WithPerRPCCredentials(insecureMetadataCreds{"authorization": "Bearer " + viper.GetString("SPICE_DB_COMMON_KEY")}),
		grpc.WithIdleTimeout(viper.GetDuration("SPICE_DB_TIMEOUT")))

	client, err := authzed.NewClient(viper.GetString("SPICE_DB_API"), options...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

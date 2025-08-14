package auth

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/pkg/rpc"
)

func New(options ...rpc.Option) (*authzed.Client, error) {
	viper.SetDefault("SPICE_DB_COMMON_KEY", "secret-shortlink-preshared-key")
	viper.SetDefault("SPICE_DB_TIMEOUT", "5s")

	config, err := rpc.SetClientConfig(options...)
	if err != nil {
		return nil, &ConfigurationError{Cause: err}
	}

	dialOptions := config.GetOptions()
	dialOptions = append(dialOptions,
		grpc.WithPerRPCCredentials(insecureMetadataCreds{"authorization": "Bearer " + viper.GetString("SPICE_DB_COMMON_KEY")}),
		grpc.WithIdleTimeout(viper.GetDuration("SPICE_DB_TIMEOUT")))

	client, err := authzed.NewClient(config.GetURI(), dialOptions...)
	if err != nil {
		return nil, &ClientInitError{Cause: err}
	}

	return client, nil
}

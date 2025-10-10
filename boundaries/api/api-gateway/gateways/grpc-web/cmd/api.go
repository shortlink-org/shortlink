//go:build goexperiment.dwarf5

/*
ShortLink application

API-service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	api_di "github.com/shortlink-org/shortlink/boundaries/api/api-gateway/gateways/grpc-web/di"
	"github.com/shortlink-org/shortlink/pkg/graceful_shutdown"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-api-grpc-web")

	// Init a new service
	service, cleanup, err := api_di.InitializeAPIService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Log.Error(r.(string)) //nolint:forcetypeassert,errcheck // simple type assertion
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	signal := graceful_shutdown.GracefulShutdown()

	cleanup()

	service.Log.Info("Service stopped", slog.String("signal", signal.String()))

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // exit code 143 is used to indicate graceful termination
}

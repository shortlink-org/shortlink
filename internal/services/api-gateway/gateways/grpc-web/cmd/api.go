//go:build goexperiment.arenas

/*
Shortlink application

API-service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/graceful_shutdown"
	_ "github.com/shortlink-org/shortlink/internal/pkg/i18n"
	api_di "github.com/shortlink-org/shortlink/internal/services/api-gateway/gateways/grpc-web/di"
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
			service.Log.Error(r.(string))
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	graceful_shutdown.GracefulShutdown()

	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) // nolint:gocritic // TODO: research
}

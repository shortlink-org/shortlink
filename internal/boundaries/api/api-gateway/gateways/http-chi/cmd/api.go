//go:build goexperiment.arenas

/*
Shortlink application

API-service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	api_di "github.com/shortlink-org/shortlink/internal/boundaries/api/api-gateway/gateways/http-chi/di"
	"github.com/shortlink-org/shortlink/internal/pkg/graceful_shutdown"
	_ "github.com/shortlink-org/shortlink/internal/pkg/i18n"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-api")

	// Init a new service
	service, cleanup, err := api_di.InitializeAPIService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Log.Error(r.(string)) //nolint:forcetypeassert // simple type assertion
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	graceful_shutdown.GracefulShutdown()

	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) // nolint:gocritic // TODO: research
}

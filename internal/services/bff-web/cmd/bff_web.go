/*
Shortlink application

BFF Web Service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/graceful_shutdown"
	bff_web_di "github.com/shortlink-org/shortlink/internal/services/bff-web/di"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-bff-web")

	// Init a new service
	service, cleanup, err := bff_web_di.InitializeBFFWebService()
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

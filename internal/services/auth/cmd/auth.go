/*
Shortlink application

Auth-service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/graceful_shutdown"
	auth_di "github.com/shortlink-org/shortlink/internal/services/auth/di"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-auth")

	// Init a new service
	service, cleanup, err := auth_di.InitializeAuthService()
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
	os.Exit(143) //nolint:gocritic // TODO: research
}

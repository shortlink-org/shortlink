/*
Bot application
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/graceful_shutdown"
	notify_di "github.com/shortlink-org/shortlink/internal/services/notify/di"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-notify")

	// Init a new service
	s, cleanup, err := notify_di.InitializeFullBotService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			s.Log.Error(r.(string))
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	graceful_shutdown.GracefulShutdown()

	// Stop the service gracefully.
	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // TODO: research
}

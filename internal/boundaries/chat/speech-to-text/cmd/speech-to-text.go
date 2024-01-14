/*
Chat boundary

Speech-to-text (STT) Service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/boundaries/chat/speech-to-text/di"
	"github.com/shortlink-org/shortlink/internal/pkg/graceful_shutdown"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "speech-to-text")

	// Init a new service
	service, cleanup, err := stt_di.InitializeSTTService()
	if err != nil {
		panic(err)
	}
	service.Log.Info("Service initialized")

	defer func() {
		if r := recover(); r != nil {
			service.Log.Error(r.(string)) //nolint:forcetypeassert // simple type assertion
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	graceful_shutdown.GracefulShutdown()

	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // TODO: research
}

/*
Chat boundary

Speech-to-text (STT) Service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	stt_di "github.com/shortlink-org/shortlink/boundaries/chat/speech-to-text/di"
	"github.com/shortlink-org/shortlink/pkg/graceful_shutdown"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
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
	signal := graceful_shutdown.GracefulShutdown()

	cleanup()

	service.Log.Info("Service stopped", field.Fields{
		"signal": signal.String(),
	})

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // exit code 143 is used to indicate graceful termination
}

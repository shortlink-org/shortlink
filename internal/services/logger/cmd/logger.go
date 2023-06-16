/*
Shortlink application

Logger-service
*/
package main

import (
	"os"

	"github.com/shortlink-org/shortlink/internal/pkg/handle_signal"
	logger_di "github.com/shortlink-org/shortlink/internal/services/logger/di"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-logger")

	// Init a new service
	service, cleanup, err := logger_di.InitializeLoggerService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Log.Error(r.(string))
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	handle_signal.WaitExitSignal()

	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143)
}

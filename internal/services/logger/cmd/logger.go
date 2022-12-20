/*
Shortlink application

Logger-service
*/
package main

import (
	"github.com/batazor/shortlink/internal/pkg/handle_signal"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/services/logger/di"
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
}

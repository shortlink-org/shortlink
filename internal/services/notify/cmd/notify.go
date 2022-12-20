/*
Bot application
*/
package main

import (
	"github.com/batazor/shortlink/internal/pkg/handle_signal"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/services/notify/di"
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
	handle_signal.WaitExitSignal()

	// Stop the service gracefully.
	cleanup()
}

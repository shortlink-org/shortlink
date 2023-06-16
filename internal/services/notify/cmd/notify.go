/*
Bot application
*/
package main

import (
	"os"

	"github.com/shortlink-org/shortlink/internal/pkg/handle_signal"
	notify_di "github.com/shortlink-org/shortlink/internal/services/notify/di"
	"github.com/spf13/viper"
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

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143)
}

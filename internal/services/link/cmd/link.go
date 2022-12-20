/*
Shortlink application

Link-service
*/
package main

import (
	"github.com/batazor/shortlink/internal/pkg/handle_signal"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/services/link/di"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-link")

	// Init a new service
	service, cleanup, err := link_di.InitializeLinkService()
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

/*
Shortlink application

Auth-service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/handle_signal"
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
	handle_signal.WaitExitSignal()

	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143)
}

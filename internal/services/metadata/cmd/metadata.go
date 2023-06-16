/*
Metadata application

Get information by links
*/
package main

import (
	"os"

	"github.com/shortlink-org/shortlink/internal/pkg/handle_signal"
	metadata_di "github.com/shortlink-org/shortlink/internal/services/metadata/di"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-metadata")

	// Init a new service
	service, cleanup, err := metadata_di.InitializeMetaDataService()
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

	// Stop the service gracefully.
	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143)
}

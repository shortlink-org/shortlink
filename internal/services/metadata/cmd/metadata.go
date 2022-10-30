/*
Metadata application

Get information by links
*/
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/services/metadata/di"
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

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Stop the service gracefully.
	cleanup()
}

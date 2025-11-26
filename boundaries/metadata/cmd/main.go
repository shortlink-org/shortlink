/*
Metadata application

Get information by links
*/
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/shortlink-org/go-sdk/graceful_shutdown"
	"github.com/spf13/viper"

	metadata_di "github.com/shortlink-org/shortlink/boundaries/metadata/internal/di"
)

const exitCodeGracefulTermination = 143

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-metadata")

	ctx := context.Background()

	// Init a new service
	service, cleanup, err := metadata_di.InitializeMetaDataService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Log.ErrorWithContext(ctx, r.(string)) //nolint:forcetypeassert,errcheck // simple type assertion
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	signal := graceful_shutdown.GracefulShutdown()

	cleanup()

	service.Log.InfoWithContext(ctx, "Service stopped", slog.String("signal", signal.String()))

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(exitCodeGracefulTermination) //nolint:gocritic // exit code 143 is used to indicate graceful termination
}

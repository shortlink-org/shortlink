//go:build goexperiment.dwarf5

/*
ShortLink application

BFF Link Service
*/
package main

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/go-sdk/graceful_shutdown"

	bff_web_di "github.com/shortlink-org/shortlink/boundaries/link/bff/internal/di"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/pkg"
)

var (
	// Build information injected at compile time
	version   = "dev"
	commit    = "none"
	buildTime = "unknown"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-bff-link")

	// Set build info metrics as per ADR-0014
	pkg.SetBuildInfo(version, commit, buildTime)

	// Init a new service
	service, cleanup, err := bff_web_di.InitializeBFFWebService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}
	service.Log.Info("Service initialized", 
		slog.String("version", version), 
		slog.String("commit", commit), 
		slog.String("build_time", buildTime))

	defer func() {
		if r := recover(); r != nil {
			service.Log.Error(r.(string)) //nolint:forcetypeassert,errcheck // simple type assertion
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	signal := graceful_shutdown.GracefulShutdown()

	cleanup()

	service.Log.Info("Service stopped", slog.String("signal", signal.String()))

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // exit code 143 is used to indicate graceful termination
}

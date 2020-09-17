/*
Metadata application

Get information by links
*/
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/batazor/shortlink/internal/config"
	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/error/status"
	"github.com/batazor/shortlink/internal/metadata/infrastructure/rpc"
)

func init() {
	// Read ENV variables
	if err := config.Init(); err != nil {
		fmt.Println(err.Error())
		os.Exit(status.ERROR_CONFIG)
	}
}

func main() {
	// Create a new context
	ctx := context.Background()

	// Init a new service
	service, cleanup, err := di.InitializeMetadataService(ctx)
	if err != nil {
		panic(err)
	}

	// Monitoring endpoints
	go http.ListenAndServe("0.0.0.0:9091", service.Monitoring) // nolint errcheck

	// Run API server
	_, err = rpc.New(service.ServerRPC, service.MetaStore, service.Log)
	if err != nil {
		panic(err)
	}

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Stop the service gracefully.
	cleanup()
}

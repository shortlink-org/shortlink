/*
Metadata application

Get information by links
*/
package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "metadata")

	// Init a new service
	service, cleanup, err := di.InitializeMetadataService()
	if err != nil { // TODO: use as helpers
		var typeErr *net.OpError
		if errors.As(err, &typeErr) {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT environment", typeErr.Addr.String()))
		}

		panic(err)
	}

	// Monitoring endpoints
	go http.ListenAndServe("0.0.0.0:9090", service.Monitoring) // nolint errcheck

	// Run API server
	_, err = metadata_rpc.New(service.ServerRPC, service.MetaStore, service.Log)
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

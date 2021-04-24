/*
Shortlink application

Link-service
*/
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/di"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "link")

	// Init a new service
	service, cleanup, err := di.InitializeLinkService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	// Run API server
	_, err = link_rpc.New(service.ServerRPC, service.LinkStore, service.Log)
	if err != nil {
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
	// close DB
	if errStoreClose := service.DB.Store.Close(); errStoreClose != nil {
		service.Log.Error(errStoreClose.Error())
	}

	cleanup()
}

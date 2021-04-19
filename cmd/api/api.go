/*
Shortlink application

API-service
*/
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/pkg/api"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "api")

	// Init a new service
	s, cleanup, err := di.InitializeAPIService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	// Run API server
	var API api.Server
	API.RunAPIServer(s.Ctx, s.Log, s.Tracer, s.ServerRPC, s.ClientRPC)

	defer func() {
		if r := recover(); r != nil {
			s.Log.Error(r.(string))
		}
	}()

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Stop the service gracefully.
	// close DB
	if err := s.DB.Store.Close(); err != nil {
		s.Log.Error(err.Error())
	}

	cleanup()
}

/*
Shortlink application

API-service
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
	"github.com/batazor/shortlink/pkg/api"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "api")

	// Init a new service
	s, cleanup, err := di.InitializeAPIService()
	if err != nil { // TODO: use as helpers
		var typeErr *net.OpError
		if errors.As(err, &typeErr) {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT environment", typeErr.Addr.String()))
		}

		panic(err)
	}

	// Monitoring endpoints
	go http.ListenAndServe("0.0.0.0:9090", s.Monitoring) // nolint errcheck

	var profiling *http.ServeMux = s.PprofEndpoint
	go http.ListenAndServe("0.0.0.0:7071", profiling) // nolint errcheck

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

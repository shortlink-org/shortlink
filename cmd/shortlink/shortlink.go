package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/pkg/api"
)

func init() {
	// Read ENV variables
	viper.AutomaticEnv()
}

func main() {
	// Create a new context
	ctx := context.Background()

	// Init a new service
	s, cleanup, err := di.InitializeFullService(ctx)
	if err != nil {
		panic(err)
	}

	// Monitoring endpoints
	go http.ListenAndServe("0.0.0.0:9090", s.Monitoring) // nolint errcheck
	var profiling *http.ServeMux
	profiling = s.PprofEndpoint
	go http.ListenAndServe("0.0.0.0:7071", profiling) // nolint errcheck

	// Run API server
	var API api.Server
	API.RunAPIServer(ctx, s.Log, s.Tracer)

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
	if err := s.DB.Close(); err != nil {
		s.Log.Error(err.Error())
	}

	cleanup()
}

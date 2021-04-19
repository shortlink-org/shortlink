/*
Logger application
*/
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/services/logger/service"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "logger")

	// Init a new service
	service, cleanup, err := di.InitializeLoggerService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	// Run logger
	logger := logger_service.Logger{
		MQ:  service.MQ,
		Log: service.Log,
	}
	logger.Use(service.Ctx)

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

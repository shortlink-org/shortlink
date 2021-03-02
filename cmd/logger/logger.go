/*
Logger application
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/pkg/mq/query"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "logger")

	// Init a new service
	service, cleanup, err := di.InitializeLoggerService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	// Monitoring endpoints
	go http.ListenAndServe("0.0.0.0:9090", service.Monitoring) // nolint errcheck

	getEventNewLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if service.MQ != nil {
			if err := service.MQ.Subscribe(getEventNewLink); err != nil {
				service.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan
			service.Log.Info(fmt.Sprintf("GET: %s", string(msg.Body)))
			msg.Context.Done()
		}
	}()

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

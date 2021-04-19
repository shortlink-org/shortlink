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

	"github.com/batazor/shortlink/internal/services/api/domain/link"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

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
			if err := service.MQ.Subscribe("shortlink", getEventNewLink); err != nil {
				service.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan

			// Convert: []byte to link.Link
			myLink := &link.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				service.Log.Error(fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				msg.Context.Done()
				continue
			}

			service.Log.InfoWithContext(msg.Context, fmt.Sprintf("GET URL: %s", myLink.Url))
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

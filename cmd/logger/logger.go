/*
Logger application
*/
package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/batazor/shortlink/internal/config"
	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/error/status"
	"github.com/batazor/shortlink/internal/mq/query"
)

func init() {
	// Read ENV variables
	if err := config.Init(); err != nil {
		fmt.Println(err.Error())
		os.Exit(status.ERROR_CONFIG)
	}
}

func main() {
	// Init a new service
	s, cleanup, err := di.InitializeLoggerService()
	if err != nil { // TODO: use as helpers
		var typeErr *net.OpError
		if errors.As(err, &typeErr) {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT environment", typeErr.Addr.String()))
		}

		panic(err)
	}

	getEventNewLink := query.Response{
		Chan: make(chan []byte),
	}

	go func() {
		if s.MQ != nil {
			if err := s.MQ.Subscribe(getEventNewLink); err != nil {
				s.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			s.Log.Info(fmt.Sprintf("GET: %s", string(<-getEventNewLink.Chan)))
		}
	}()

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
	cleanup()
}

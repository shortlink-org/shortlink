package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/mq/query"
	"github.com/batazor/shortlink/internal/transform"
	"github.com/batazor/shortlink/pkg/link"
)

func init() {
	// Read ENV variables
	viper.AutomaticEnv()
}

func main() {
	// Create a new context
	ctx := context.Background()

	// Init a new service
	s, cleanup, err := di.InitializeLoggerService(ctx)
	if err != nil {
		panic(err)
	}

	test := query.Response{
		Chan: make(chan []byte),
	}

	go func() {
		if s.MQ != nil {
			if err := s.MQ.Subscribe(test); err != nil {
				s.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			msg := <-test.Chan
			response, err := transform.Deserialize(msg, &link.Link{})
			if err != nil {
				s.Log.Error(err.Error())
				continue
			}

			link := response.(*link.Link)

			s.Log.Info(fmt.Sprintf("GET Url: %s", link.Url))
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

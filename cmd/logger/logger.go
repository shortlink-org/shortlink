package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/di"
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

	test := make(chan []byte)

	go func() {
		if s.MQ != nil {
			if err := s.MQ.Subscribe(test); err != nil {
				s.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			s.Log.Info(fmt.Sprintf("GET: %s", string(<-test)))
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
	// flushes buffer, if any
	s.Log.Close()
}

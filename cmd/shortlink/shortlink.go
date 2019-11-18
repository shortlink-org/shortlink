package main

import (
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// Read ENV variables
	viper.AutomaticEnv()
}

func main() {
	// Init a new service
	s := Service{}
	go s.Start()

	defer func() {
		if r := recover(); r != nil {
			s.log.Error(r.(string))
		}
	}()

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Stop the service gracefully.
	s.Stop()
}

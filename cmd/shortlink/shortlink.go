package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Init a new service
	s := Service{}
	go s.Start()

	defer func() {
		if r := recover(); r != nil {
			s.log.Error(r.(string))
		}
	}()

	// Test Event
	// TODO: Delete next line
	s.tracer.StartSpan("test").Finish()

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Stop the service gracefully.
	s.Stop()
}

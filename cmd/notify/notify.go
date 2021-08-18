/*
Bot application
*/
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/services/notify/application"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "notify")

	// Init a new service
	s, cleanup, err := di.InitializeNotifyService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	// Run bot
	bot, err := application.New(s.MQ, s.Log)
	if err != nil { // TODO: use as helpers
		panic(err)
	}
	bot.Use(s.Ctx)

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

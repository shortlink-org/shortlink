/*
Bot application
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/bot"
	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/mq/query"
	"github.com/batazor/shortlink/internal/notify"
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
	s, cleanup, err := di.InitializeBotService(ctx)
	if err != nil {
		panic(err)
	}

	getEventNewLink := query.Response{
		Chan: make(chan []byte),
	}

	// Run bot
	b := bot.Bot{}
	b.Use(ctx)

	go func() {
		if s.MQ != nil {
			if err := s.MQ.Subscribe(getEventNewLink); err != nil {
				s.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan

			// []byte to link.Link
			myLink := &link.Link{}
			if err := proto.Unmarshal(msg, myLink); err != nil {
				s.Log.Error(fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				continue
			}

			notify.Publish(ctx, bot.METHOD_NEW_LINK, myLink, nil, "")
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

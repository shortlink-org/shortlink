/*
Bot application
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain/link"
	"github.com/batazor/shortlink/internal/services/bot/service"
	bot_type "github.com/batazor/shortlink/internal/services/bot/type"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "bot")

	// Init a new service
	s, cleanup, err := di.InitializeBotService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	// Monitoring endpoints
	go http.ListenAndServe("0.0.0.0:9090", s.Monitoring) // nolint errcheck

	getEventNewLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	// Run bot
	b := service.Bot{}
	b.Use(s.Ctx)

	go func() {
		if s.MQ != nil {
			if err := s.MQ.Subscribe("shortlink", getEventNewLink); err != nil {
				s.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan

			// Convert: []byte to link.Link
			myLink := &link.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				s.Log.ErrorWithContext(msg.Context, fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				continue
			}

			s.Log.InfoWithContext(msg.Context, "Get new LINK", field.Fields{"url": myLink.Url})
			notify.Publish(s.Ctx, bot_type.METHOD_NEW_LINK, myLink, nil)
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

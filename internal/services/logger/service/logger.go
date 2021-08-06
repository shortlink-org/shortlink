package logger_service

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	bot_type "github.com/batazor/shortlink/internal/services/notify/type"
)

func (l *Logger) Use(_ context.Context) { // nolint unused
	// TODO: refactoring this code
	getEventNewLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if l.MQ != nil {
			if err := l.MQ.Subscribe("shortlink", getEventNewLink); err != nil {
				l.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan

			// Convert: []byte to link.Link
			myLink := &v1.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				l.Log.Error(fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				msg.Context.Done()
				continue
			}

			l.Log.InfoWithContext(msg.Context, fmt.Sprintf("GET URL: %s", myLink.Url))
			msg.Context.Done()
		}
	}()
}

// Notify ...
func (l *Logger) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case bot_type.METHOD_NEW_LINK:
		if addLink, ok := payload.(*v1.Link); ok {
			l.Send(ctx, addLink)
		}
	}

	return notify.Response{}
}

func (l *Logger) Send(ctx context.Context, link *v1.Link) {
	payload := fmt.Sprintf("LINK: %s", link.Url)

	notify.Publish(ctx, bot_type.METHOD_SEND_NEW_LINK, payload, nil)
}

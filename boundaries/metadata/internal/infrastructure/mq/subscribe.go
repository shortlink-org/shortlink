package metadata_mq

import (
	"context"
	"log/slog"

	linkrpc "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/infrastructure/rpc/link/v1"
	"github.com/segmentio/encoding/json"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/mq/query"
)

const linkCreatedEvent = "shortlink.link.event.created"

// SubscribeLinkCreated subscribes to link creation events from Kafka
// When a link is created, it processes the link URL to extract metadata
func (e *Event) SubscribeLinkCreated(log logger.Logger) error {
	getCreatedLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if err := e.mq.Subscribe(context.Background(), linkCreatedEvent, getCreatedLink); err != nil {
			log.ErrorWithContext(context.Background(), "failed to subscribe to link created events",
				slog.String("error", err.Error()),
				slog.String("event", linkCreatedEvent),
			)
		}
	}()

	go func() {
		for {
			msg := <-getCreatedLink.Chan

			// Convert: []byte to link.Link
			myLink := &linkrpc.Link{}
			if err := json.Unmarshal(msg.Body, myLink); err != nil {
				log.ErrorWithContext(msg.Context, "Error unmarshaling link created event",
					slog.String("error", err.Error()),
				)
				msg.Context.Done()
				continue
			}

			// Get URL from link
			linkURL := myLink.GetUrl()
			if linkURL == "" {
				log.ErrorWithContext(msg.Context, "Link URL is nil")
				msg.Context.Done()
				continue
			}

			// Process metadata for the link URL
			_, err := e.metadataUC.Add(msg.Context, linkURL)
			if err != nil {
				log.ErrorWithContext(msg.Context, "Error processing metadata for link",
					slog.String("error", err.Error()),
					slog.String("url", linkURL),
					slog.String("hash", myLink.GetHash()),
				)
			} else {
				log.InfoWithContext(msg.Context, "Successfully processed metadata for link",
					slog.String("url", linkURL),
					slog.String("hash", myLink.GetHash()),
				)
			}

			msg.Context.Done()
		}
	}()

	return nil
}

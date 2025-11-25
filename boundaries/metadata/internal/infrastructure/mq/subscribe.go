package metadata_mq

import (
	"context"
	"log/slog"

	linkrpc "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/infrastructure/rpc/link/v1"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/segmentio/encoding/json"

	"github.com/shortlink-org/go-sdk/logger"
)

const linkCreatedEvent = "shortlink.link.event.created"

// SubscribeLinkCreated subscribes to link creation events from Kafka
// When a link is created, it processes the link URL to extract metadata
func (e *Event) SubscribeLinkCreated(ctx context.Context, log logger.Logger) error {
	messages, err := e.subscriber.Subscribe(ctx, linkCreatedEvent)
	if err != nil {
		log.ErrorWithContext(ctx, "failed to subscribe to link created events",
			slog.String("error", err.Error()),
			slog.String("event", linkCreatedEvent),
		)
		return err
	}

	go func() {
		for msg := range messages {
			// Convert: []byte to link.Link
			myLink := &linkrpc.Link{}
			if err := json.Unmarshal(msg.Payload, myLink); err != nil {
				log.ErrorWithContext(msg.Context(), "Error unmarshaling link created event",
					slog.String("error", err.Error()),
				)
				msg.Nack()
				continue
			}

			// Get URL from link
			linkURL := myLink.GetUrl()
			if linkURL == "" {
				log.ErrorWithContext(msg.Context(), "Link URL is nil")
				msg.Nack()
				continue
			}

			// Process metadata for the link URL
			_, err := e.metadataUC.Add(msg.Context(), linkURL)
			if err != nil {
				log.ErrorWithContext(msg.Context(), "Error processing metadata for link",
					slog.String("error", err.Error()),
					slog.String("url", linkURL),
					slog.String("hash", myLink.GetHash()),
				)
				msg.Nack()
			} else {
				log.InfoWithContext(msg.Context(), "Successfully processed metadata for link",
					slog.String("url", linkURL),
					slog.String("hash", myLink.GetHash()),
				)
				msg.Ack()
			}
		}
	}()

	return nil
}

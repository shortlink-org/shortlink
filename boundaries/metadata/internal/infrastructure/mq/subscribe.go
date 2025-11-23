package metadata_mq

import (
	"context"
	"log/slog"

	"github.com/segmentio/encoding/json"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/mq/query"
	link_domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// SubscribeLinkCreated subscribes to link creation events from Kafka
// When a link is created, it processes the link URL to extract metadata
func (e *Event) SubscribeLinkCreated(log logger.Logger) error {
	getCreatedLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if err := e.mq.Subscribe(context.Background(), link_domain.MQ_EVENT_LINK_CREATED, getCreatedLink); err != nil {
			log.ErrorWithContext(context.Background(), "failed to subscribe to link created events",
				slog.String("error", err.Error()),
				slog.String("event", link_domain.MQ_EVENT_LINK_CREATED),
			)
		}
	}()

	go func() {
		for {
			msg := <-getCreatedLink.Chan

			// Convert: []byte to link.Link
			myLink := &link_domain.Link{}
			if err := json.Unmarshal(msg.Body, myLink); err != nil {
				log.ErrorWithContext(msg.Context, "Error unmarshaling link created event",
					slog.String("error", err.Error()),
				)
				msg.Context.Done()
				continue
			}

			// Get URL from link
			url := myLink.GetUrl()
			if url == nil {
				log.ErrorWithContext(msg.Context, "Link URL is nil")
				msg.Context.Done()
				continue
			}

			// Convert Url to string - assuming Url has a String() method or similar
			// If not, we may need to access the underlying value
			linkURL := url.String()

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


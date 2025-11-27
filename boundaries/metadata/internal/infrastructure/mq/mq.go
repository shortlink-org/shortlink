/*
MQ Endpoint
*/

package metadata_mq

import (
	"go.opentelemetry.io/otel/trace"

	"github.com/ThreeDotsLabs/watermill/message"

	metadata_uc "github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/metadata"
)

type Event struct {
	subscriber message.Subscriber
	metadataUC *metadata_uc.UC
	tracer     trace.TracerProvider
}

func New(subscriber message.Subscriber, metadataUC *metadata_uc.UC, tracer trace.TracerProvider) (*Event, error) {
	return &Event{
		subscriber: subscriber,
		metadataUC: metadataUC,
		tracer:     tracer,
	}, nil
}

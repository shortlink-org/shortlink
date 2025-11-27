/*
MQ Endpoint
*/

package metadata_mq

import (
	"github.com/ThreeDotsLabs/watermill/message"

	metadata_uc "github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/metadata"
)

type Event struct {
	subscriber message.Subscriber
	metadataUC *metadata_uc.UC
}

func New(subscriber message.Subscriber, metadataUC *metadata_uc.UC) (*Event, error) {
	return &Event{
		subscriber: subscriber,
		metadataUC: metadataUC,
	}, nil
}

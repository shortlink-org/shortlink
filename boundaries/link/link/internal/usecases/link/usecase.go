/*
Link UC. Application layer
*/
package link

import (
	"github.com/authzed/authzed-go/v1"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/notify"
)

type UC struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[domain.Link]

	// Common
	log logger.Logger

	// Security
	permission *authzed.Client

	// Delivery
	mq mq.MQ
	// MetadataClient metadata_rpc.MetadataServiceClient

	// Repository
	store crud.Repository
}

// New creates a new link usecase
func New(log logger.Logger, dataBus mq.MQ, metadataService any, store *crud.Store, permissionClient *authzed.Client) (*UC, error) {
	service := &UC{
		log: log,

		// Security
		permission: permissionClient,

		// Delivery
		mq: dataBus,
		// MetadataClient: metadataService,

		// Repository
		store: store,
	}

	return service, nil
}

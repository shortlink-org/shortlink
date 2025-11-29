/*
Link UC. Application layer
*/
package link

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/shortlink-org/go-sdk/cqrs/bus"
	"github.com/shortlink-org/go-sdk/logger"

	"github.com/shortlink-org/go-sdk/kratos"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud"
)

type UC struct {
	// Common
	log logger.Logger

	// Security
	permission *authzed.Client
	kratos     kratos.KratosClient

	// Delivery
	eventBus *bus.EventBus // CQRS EventBus for publishing events

	// MetadataClient metadata_rpc.MetadataServiceClient

	// Repository
	store crud.Repository
}

// New creates a new link usecase
func New(
	log logger.Logger,
	metadataService any,
	store crud.Repository,
	permissionClient *authzed.Client,
	kratosClient kratos.KratosClient,
	eventBus *bus.EventBus,
) (*UC, error) {
	service := &UC{
		log: log,

		// Security
		permission: permissionClient,
		kratos:     kratosClient,

		// Delivery
		eventBus: eventBus,

		// MetadataClient: metadataService,

		// Repository
		store: store,
	}

	return service, nil
}

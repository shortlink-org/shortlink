/*
OMS UC. Application layer
*/
package cart

import (
	"github.com/authzed/authzed-go/v1"
	"go.temporal.io/sdk/client"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

type UC struct {
	// Common
	log logger.Logger

	// Security
	permission *authzed.Client

	// Temporal
	temporalClient client.Client
}

// New creates a new cart usecase
func New(log logger.Logger, permissionClient *authzed.Client, temporalClient client.Client) (*UC, error) {
	service := &UC{
		log: log,

		// Security
		permission: permissionClient,

		// Temporal
		temporalClient: temporalClient,
	}

	return service, nil
}

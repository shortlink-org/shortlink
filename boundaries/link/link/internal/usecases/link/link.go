/*
Link UC. Application layer
*/
package link

import (
	"context"
	"fmt"

	"github.com/authzed/authzed-go/v1"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/notify"
)

type UC struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[domain.Link]

	// Security
	permission *authzed.Client

	// Delivery
	mq mq.MQ
	// MetadataClient metadata_rpc.MetadataServiceClient

	// Repository
	store crud.Repository

	log logger.Logger
}

func New(log logger.Logger, dataBus mq.MQ, metadataService any, store crud.Repository, permissionClient *authzed.Client) (*UC, error) {
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

func errorHelper(ctx context.Context, log logger.Logger, errs []error) error {
	if len(errs) > 0 {
		errList := field.Fields{}
		for index := range errs {
			errList[fmt.Sprintf("stack error: %d", index)] = errs[index]
		}

		log.ErrorWithContext(ctx, "Error create a new link", errList)

		return ErrCreateLink
	}

	return nil
}

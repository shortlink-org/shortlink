/*
Link Service. Application layer
*/
package link

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/cqrs/query"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
)

type Service struct {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event

	// Delivery
	MetadataClient metadata_rpc.MetadataClient

	// Repository
	cqsStore   *store.Store
	queryStore *query.Store

	logger logger.Logger
}

func New(logger logger.Logger, metadataService metadata_rpc.MetadataClient, cqsStore *store.Store, queryStore *query.Store) (*Service, error) {
	service := &Service{
		MetadataClient: metadataService,
		cqsStore:       cqsStore,
		queryStore:     queryStore,
		logger:         logger,
	}

	// Subscribe to event
	service.EventHandler()

	return service, nil
}

func errorHelper(ctx context.Context, logger logger.Logger, errs []error) error {
	if len(errs) > 0 {
		errList := field.Fields{}
		for index := range errs {
			errList[fmt.Sprintf("stack error: %d", index)] = errs[index]
		}

		logger.ErrorWithContext(ctx, "Error create a new link", errList)
		return fmt.Errorf("Error create a new link")
	}

	return nil
}

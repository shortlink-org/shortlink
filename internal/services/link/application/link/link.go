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
	link_store "github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
)

type Service struct {
	// system event
	notify.Subscriber // Observer interface for subscribe on system event

	// Delivery
	MetadataClient metadata_rpc.MetadataClient

	// Repository
	*link_store.LinkStore

	logger logger.Logger
}

func New(logger logger.Logger, metadataService metadata_rpc.MetadataClient, linkStore *link_store.LinkStore) (*Service, error) {
	service := &Service{
		MetadataClient: metadataService,
		LinkStore:      linkStore,
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

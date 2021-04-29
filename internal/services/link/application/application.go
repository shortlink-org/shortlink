/*
Metadata Service. Application layer
*/
package application

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	link_store "github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
	"github.com/batazor/shortlink/pkg/saga"
)

type Service struct {
	// Delivery
	MetadataClient metadata_rpc.MetadataClient

	// Repository
	*link_store.LinkStore

	logger logger.Logger
}

func New(logger logger.Logger, metadataService metadata_rpc.MetadataClient, linkStore *link_store.LinkStore) (*Service, error) {
	return &Service{
		MetadataClient: metadataService,
		LinkStore:      linkStore,
		logger:         logger,
	}, nil
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

func (s *Service) AddLink(ctx context.Context, link *link.Link) (*link.Link, error) {
	const (
		SAGA_NAME              = "ADD_LINK"
		SAGA_STEP_STORE_SAVE   = "SAGA_STEP_STORE_SAVE"
		SAGA_STEP_METADATA_GET = "SAGA_STEP_METADATA_GET"
	)

	// create a new saga for create a new link
	sagaAddLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: save to Store
	_, errs = sagaAddLink.AddStep(SAGA_STEP_STORE_SAVE).
		Then(func(ctx context.Context) error {
			var err error
			link, err = s.Store.Add(ctx, link)
			return err
		}).
		Reject(func(ctx context.Context) error {
			return s.Store.Delete(ctx, link.Hash)
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: request to metadata
	_, errs = sagaAddLink.AddStep(SAGA_STEP_METADATA_GET).
		Then(func(ctx context.Context) error {
			_, err := s.MetadataClient.Set(ctx, &metadata_rpc.SetMetaRequest{
				Id: link.Url,
			})
			return err
		}).
		Reject(func(ctx context.Context) error {
			return s.Store.Delete(ctx, link.Hash)
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaAddLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return link, nil
}

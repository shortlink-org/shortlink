//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/internal/pkg/notify"
	metadata "github.com/batazor/shortlink/internal/services/metadata/application"
	metadata_domain "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"
	api_mq "github.com/batazor/shortlink/internal/services/metadata/infrastructure/mq"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	meta_store "github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

type MetaDataService struct {
	Logger logger.Logger

	// Delivery
	linkMQ            *api_mq.Event
	metadataRPCServer *metadata_rpc.Metadata

	// Application
	service *metadata.Service

	// Repository
	metadataStore *meta_store.MetaStore
}

// MetaDataService =====================================================================================================
var MetaDataSet = wire.NewSet(
	// Delivery
	InitLinkMQ,
	NewMetaDataRPCServer,

	// applications
	NewMetaDataApplication,

	// repository
	NewMetaDataStore,

	NewMetaDataService,
)

func InitLinkMQ(ctx context.Context, log logger.Logger, mq mq.MQ) (*api_mq.Event, error) {
	linkMQ := &api_mq.Event{
		MQ: mq,
	}

	// Subscribe to Event
	notify.Subscribe(metadata_domain.METHOD_ADD, linkMQ)

	return linkMQ, nil
}

func NewMetaDataStore(ctx context.Context, logger logger.Logger, db *db.Store) (*meta_store.MetaStore, error) {
	store := &meta_store.MetaStore{}
	metadataStore, err := store.Use(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	return metadataStore, nil
}

func NewMetaDataApplication(store *meta_store.MetaStore) (*metadata.Service, error) {
	metadataService, err := metadata.New(store)
	if err != nil {
		return nil, err
	}

	return metadataService, nil
}

func NewMetaDataRPCServer(runRPCServer *rpc.RPCServer, application *metadata.Service, log logger.Logger) (*metadata_rpc.Metadata, error) {
	metadataRPCServer, err := metadata_rpc.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return metadataRPCServer, nil
}

func NewMetaDataService(
	log logger.Logger,

	// Application
	service *metadata.Service,

	// Delivery
	linkMQ *api_mq.Event,
	metadataRPCServer *metadata_rpc.Metadata,

	// Repository
	metadataStore *meta_store.MetaStore,
) (*MetaDataService, error) {
	return &MetaDataService{
		Logger: log,

		// Application
		service: service,

		// Delivery
		linkMQ:            linkMQ,
		metadataRPCServer: metadataRPCServer,

		// Repository
		metadataStore: metadataStore,
	}, nil
}

func InitializeMetaDataService(ctx context.Context, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq mq.MQ) (*MetaDataService, func(), error) {
	panic(wire.Build(MetaDataSet))
}

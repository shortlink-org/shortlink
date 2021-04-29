//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	metadata "github.com/batazor/shortlink/internal/services/metadata/application"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
	meta_store "github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

type MetaDataService struct {
	Logger logger.Logger

	// Delivery
	metadataRPCServer *metadata_rpc.Metadata

	// Application
	service *metadata.Service

	// Repository
	metadataStore *meta_store.MetaStore
}

// MetaDataService =====================================================================================================
var MetaDataSet = wire.NewSet(
	// gRPC server
	NewMetaDataRPCServer,

	// applications
	NewMetaDataApplication,

	// repository
	NewMetaDataStore,

	NewMetaDataService,
)

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
	metadataRPCServer *metadata_rpc.Metadata,
	metadataStore *meta_store.MetaStore,
	service *metadata.Service,
) (*MetaDataService, error) {
	return &MetaDataService{
		Logger:            log,
		metadataRPCServer: metadataRPCServer,
		metadataStore:     metadataStore,
		service:           service,
	}, nil
}

func InitializeMetaDataService(ctx context.Context, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store) (*MetaDataService, func(), error) {
	panic(wire.Build(MetaDataSet))
}

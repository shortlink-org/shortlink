//go:generate wire
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
MetaData Service DI-package
*/
package metadata_di

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	mq_di "github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/store"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	metadata "github.com/shortlink-org/shortlink/internal/services/metadata/application"
	metadata_domain "github.com/shortlink-org/shortlink/internal/services/metadata/domain/metadata/v1"
	metadata_mq "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/mq"
	metadata_rpc "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	meta_store "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type MetaDataService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        *trace.TracerProvider
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	metadataMQ        *metadata_mq.Event
	metadataRPCServer *metadata_rpc.Metadata

	// Application
	service *metadata.Service

	// Repository
	metadataStore *meta_store.MetaStore
}

// MetaDataService =====================================================================================================
var MetaDataSet = wire.NewSet(
	di.DefaultSet,
	mq_di.New,
	store.New,
	rpc.InitServer,

	// Delivery
	InitMetadataMQ,
	NewMetaDataRPCServer,

	// Applications
	NewMetaDataApplication,

	// repository
	NewMetaDataStore,

	NewMetaDataService,
)

func InitMetadataMQ(ctx context.Context, log logger.Logger, mq *v1.DataBus) (*metadata_mq.Event, error) {
	metadataMQ, err := metadata_mq.New(mq)
	if err != nil {
		return nil, err
	}

	// Subscribe to Event
	notify.Subscribe(metadata_domain.METHOD_ADD, metadataMQ)

	return metadataMQ, nil
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
	// Common
	log logger.Logger,
	config *config.Config,

	// Observability
	monitoring *http.ServeMux,
	tracer *trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Application
	service *metadata.Service,

	// Delivery
	metadataMQ *metadata_mq.Event,
	metadataRPCServer *metadata_rpc.Metadata,

	// Repository
	metadataStore *meta_store.MetaStore,
) (*MetaDataService, error) {
	return &MetaDataService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		// Application
		service: service,

		// Delivery
		metadataMQ:        metadataMQ,
		metadataRPCServer: metadataRPCServer,

		// Repository
		metadataStore: metadataStore,
	}, nil
}

func InitializeMetaDataService() (*MetaDataService, func(), error) {
	panic(wire.Build(MetaDataSet))
}

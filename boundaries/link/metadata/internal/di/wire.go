//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
MetaData UC DI-package
*/
package metadata_di

import (
	"context"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	metadata_domain "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/domain/metadata/v1"
	metadata_mq "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/mq"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/repository/media"
	meta_store "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/repository/store"
	metadata_rpc "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/infrastructure/rpc/metadata/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/usecases/metadata"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/usecases/parsers"
	"github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/usecases/screenshot"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	mq_di "github.com/shortlink-org/shortlink/pkg/di/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/store"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/notify"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
	"github.com/shortlink-org/shortlink/pkg/s3"
)

type MetaDataService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint

	// Delivery
	metadataMQ        *metadata_mq.Event
	metadataRPCServer *metadata_rpc.Metadata

	// Application
	service *parsers.UC

	// Repository
	metadataStore *meta_store.MetaStore
}

// MetaDataService =====================================================================================================
var MetaDataSet = wire.NewSet(
	di.DefaultSet,
	mq_di.New,
	store.New,
	rpc.InitServer,
	s3.New,

	// Delivery
	InitMetadataMQ,
	NewMetaDataRPCServer,

	// Applications
	NewParserUC,
	NewScreenshotUC,
	NewMetadataUC,

	// repository
	NewMetaDataStore,
	NewMetaDataMediaStore,

	NewMetaDataService,
)

func InitMetadataMQ(ctx context.Context, dataBus mq.MQ) (*metadata_mq.Event, error) {
	metadataMQ, err := metadata_mq.New(dataBus)
	if err != nil {
		return nil, err
	}

	// Subscribe to Event
	notify.Subscribe(metadata_domain.METHOD_ADD, metadataMQ)

	return metadataMQ, nil
}

func NewMetaDataStore(ctx context.Context, log logger.Logger, db db.DB) (*meta_store.MetaStore, error) {
	store := &meta_store.MetaStore{}
	metadataStore, err := store.Use(ctx, log, db)
	if err != nil {
		return nil, err
	}

	return metadataStore, nil
}

func NewMetaDataMediaStore(ctx context.Context, s3 *s3.Client) (*s3Repository.Service, error) {
	client, err := s3Repository.New(ctx, s3)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewParserUC(store *meta_store.MetaStore) (*parsers.UC, error) {
	metadataService, err := parsers.New(store)
	if err != nil {
		return nil, err
	}

	return metadataService, nil
}

func NewScreenshotUC(ctx context.Context, media *s3Repository.Service) (*screenshot.UC, error) {
	metadataService, err := screenshot.New(ctx, media)
	if err != nil {
		return nil, err
	}

	return metadataService, nil
}

func NewMetadataUC(log logger.Logger, parsersUC *parsers.UC, screenshotUC *screenshot.UC) (*metadata.UC, error) {
	metadataService, err := metadata.New(log, parsersUC, screenshotUC)
	if err != nil {
		return nil, err
	}

	return metadataService, nil
}

func NewMetaDataRPCServer(log logger.Logger, runRPCServer *rpc.Server, parsersUC *parsers.UC, screenshotUC *screenshot.UC, metadataUC *metadata.UC) (*metadata_rpc.Metadata, error) {
	metadataRPCServer, err := metadata_rpc.New(log, runRPCServer, parsersUC, screenshotUC, metadataUC)
	if err != nil {
		return nil, err
	}

	return metadataRPCServer, nil
}

func NewMetaDataService(
	// Common
	log logger.Logger,
	config *config.Config,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	// Application
	service *parsers.UC,

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

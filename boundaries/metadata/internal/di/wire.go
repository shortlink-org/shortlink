//go:generate go tool wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
MetaData UC DI-package
*/
package metadata_di

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
	"github.com/shortlink-org/go-sdk/auth/permission"
	"github.com/shortlink-org/go-sdk/cache"
	"github.com/shortlink-org/go-sdk/config"
	shortctx "github.com/shortlink-org/go-sdk/context"
	"github.com/shortlink-org/go-sdk/db"
	"github.com/shortlink-org/go-sdk/flags"
	"github.com/shortlink-org/go-sdk/flight_trace"
	rpc "github.com/shortlink-org/go-sdk/grpc"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/watermill"
	watermill_kafka "github.com/shortlink-org/go-sdk/watermill/backends/kafka"
	"github.com/shortlink-org/go-sdk/observability/metrics"
	"github.com/shortlink-org/go-sdk/observability/profiling"
	"github.com/shortlink-org/go-sdk/observability/tracing"
	"github.com/shortlink-org/go-sdk/s3"
	"go.opentelemetry.io/otel/metric"
	api "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/go-sdk/cqrs/bus"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"

	cqrs_registry "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/cqrs"
	metadata_mq "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/mq"
	s3Repository "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/repository/media"
	meta_store "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/repository/store"
	metadata_rpc "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/rpc/metadata/v1"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/metadata"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/parsers"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/screenshot"
)

type MetaDataService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Metrics       *metrics.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	FlightTrace   *flight_trace.Recorder

	// Delivery
	metadataMQ        *metadata_mq.Event
	metadataRPCServer *metadata_rpc.Metadata

	// Application
	service *parsers.UC

	// Repository
	metadataStore *meta_store.MetaStore
}

// DefaultSet ==========================================================================================================
var DefaultSet = wire.NewSet(
	shortctx.New,
	flags.New,
	config.New,
	logger.NewDefault,
	tracing.New,
	metrics.New,
	cache.New,
	profiling.New,
	flight_trace.New,
)

// CQRSSet =============================================================================================================
// CQRS wire set for event-driven architecture components
// Provides: EventRegistry, ShortlinkNamer, ProtoMarshaler, EventBus, CommandBus
// Requires: message.Publisher (from watermill.Client)
// Note: ShortlinkNamer is a singleton to ensure consistent naming across all components
var CQRSSet = wire.NewSet(
	// CQRS Registry and Namer (singleton)
	cqrs_registry.NewEventRegistry,
	cqrs_registry.NewShortlinkNamer,
	// ProtoMarshaler (depends on namer)
	cqrs_registry.NewProtoMarshaler,
	// Bind concrete ProtoMarshaler to Marshaler interface
	wire.Bind(new(cqrsmessage.Marshaler), new(*cqrsmessage.ProtoMarshaler)),
	// EventBus and CommandBus (depend on namer and marshaler)
	cqrs_registry.NewEventBus,
	cqrs_registry.NewCommandBus,
)

// MetaDataService =====================================================================================================
var MetaDataSet = wire.NewSet(
	DefaultSet,
	permission.New,
	wire.FieldsOf(new(*metrics.Monitoring), "Prometheus", "Metrics"),
	wire.Bind(new(metric.MeterProvider), new(*api.MeterProvider)),
	db.New,
	wire.Bind(new(watermill.Backend), new(*watermill_kafka.Backend)),
	watermill_kafka.New,
	wire.Value([]watermill.Option{}),
	watermill.New,
	wire.FieldsOf(new(*watermill.Client), "Publisher", "Subscriber"),
	rpc.InitServer,
	s3.New,

	// CQRS (using CQRSSet)
	CQRSSet,

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

func InitMetadataMQ(
	ctx context.Context,
	log logger.Logger,
	subscriber message.Subscriber,
	metadataUC *metadata.UC,
	registry *bus.TypeRegistry,
	marshaler cqrsmessage.Marshaler,
) (*metadata_mq.Event, error) {
	metadataMQ, err := metadata_mq.New(subscriber, metadataUC)
	if err != nil {
		return nil, err
	}

	// Subscribe to link creation events from Kafka using TypeRegistry and ProtoMarshaler
	// TypeRegistry resolves event type, ProtoMarshaler handles deserialization
	// This eliminates manual reflect usage - only proto reflection for field access
	// Uses canonical topic name: link.link.created.v1
	if err := metadataMQ.SubscribeLinkCreated(ctx, log, registry, marshaler); err != nil {
		return nil, err
	}

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

func NewParserUC(store *meta_store.MetaStore, eventBus *bus.EventBus, log logger.Logger) (*parsers.UC, error) {
	metadataService, err := parsers.New(store, eventBus, log)
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

func NewMetadataUC(log logger.Logger, parsersUC *parsers.UC, screenshotUC *screenshot.UC, eventBus *bus.EventBus) (*metadata.UC, error) {
	metadataService, err := metadata.New(log, parsersUC, screenshotUC, eventBus)
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

	// Observability
	metrics *metrics.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	flightTrace *flight_trace.Recorder,

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
		Metrics:       metrics,
		PprofEndpoint: pprofHTTP,
		FlightTrace:   flightTrace,

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

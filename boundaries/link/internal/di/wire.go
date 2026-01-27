//go:generate go tool wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Link UC DI-package
*/
package link_di

import (
	"context"

	"github.com/authzed/authzed-go/v1"
	"github.com/google/wire"
	shortctx "github.com/shortlink-org/go-sdk/context"
	"github.com/shortlink-org/go-sdk/cqrs/bus"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"
	"github.com/shortlink-org/go-sdk/flags"
	"github.com/shortlink-org/go-sdk/flight_trace"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/observability/tracing"
	"go.opentelemetry.io/otel/metric"
	api "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	cqrs_registry "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/cqrs"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/cqrs/query"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud"
	cqrs_rpc "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/rpc/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/rpc/run"
	sitemap_rpc "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/rpc/sitemap/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/usecases/link"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/usecases/link_cqrs"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/usecases/sitemap"

	"github.com/shortlink-org/go-sdk/auth/permission"
	"github.com/shortlink-org/go-sdk/cache"
	"github.com/shortlink-org/go-sdk/config"
	"github.com/shortlink-org/go-sdk/db"
	rpc "github.com/shortlink-org/go-sdk/grpc"
	"github.com/shortlink-org/go-sdk/kratos"
	"github.com/shortlink-org/go-sdk/observability/metrics"
	"github.com/shortlink-org/go-sdk/observability/profiling"
	"github.com/shortlink-org/go-sdk/watermill"
	watermill_kafka "github.com/shortlink-org/go-sdk/watermill/backends/kafka"
)

type LinkService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Metrics       *metrics.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	FlightTrace   *flight_trace.Recorder

	// Security
	authPermission *authzed.Client

	// Delivery
	run               *run.Response
	linkRPCServer     *link_rpc.LinkRPC
	linkCQRSRPCServer *cqrs_rpc.LinkRPC
	sitemapRPCServer  *sitemap_rpc.Sitemap

	// Application
	linkService     *link.UC
	linkCQRSService *link_cqrs.Service
	sitemapService  *sitemap.Service

	// Repository
	linkStore *crud.Store

	// CQRS
	cqsStore   *cqs.Store
	queryStore *query.Store
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

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(
	// Common
	DefaultSet,
	permission.New,
	wire.FieldsOf(new(*metrics.Monitoring), "Prometheus", "Metrics"),
	wire.Bind(new(metric.MeterProvider), new(*api.MeterProvider)),
	db.New,

	// Infrastructure
	kratos.New,

	// Delivery
	wire.Bind(new(watermill.Backend), new(*watermill_kafka.Backend)),
	watermill_kafka.New,
	wire.Value([]watermill.Option{}),
	watermill.New,
	wire.FieldsOf(new(*watermill.Client), "Publisher", "Subscriber"),
	rpc.InitServer,
	NewRPCClient,
	link_rpc.New,
	cqrs_rpc.New,
	sitemap_rpc.New,
	NewRunRPCServer,

	link_rpc.NewLinkServiceClient,
	// metadata_rpc.NewMetadataServiceClient,

	// CQRS (using CQRSSet)
	CQRSSet,

	// Applications
	NewLinkApplication,
	link_cqrs.New,
	NewSitemapService,

	// repository
	crud.New,
	cqs.New,
	query.New,
	// Bind concrete Store types to Repository interfaces
	wire.Bind(new(crud.Repository), new(*crud.Store)),
	wire.Bind(new(cqs.Repository), new(*cqs.Store)),
	wire.Bind(new(query.Repository), new(*query.Store)),

	NewLinkService,
)

func NewRPCClient(
	ctx context.Context,
	log logger.Logger,
	cfg *config.Config,
	metrics *metrics.Monitoring,
	tracer trace.TracerProvider,
) (*grpc.ClientConn, func(), error) {
	// Initialize gRPC Client's interceptor.
	opts := []rpc.Option{
		rpc.WithAuthForward(),
		rpc.WithMetrics(metrics.Prometheus),
		rpc.WithTracer(tracer, metrics.Prometheus, metrics.Metrics),
		rpc.WithTimeout(),
		rpc.WithLogger(log),
	}

	runRPCClient, cleanup, err := rpc.InitClient(ctx, log, cfg, opts...)
	if err != nil {
		return nil, nil, err
	}

	return runRPCClient, cleanup, nil
}

func NewLinkApplication(
	log logger.Logger,
	eventBus *bus.EventBus,
	store crud.Repository,
	authPermission *authzed.Client,
	kratosClient *kratos.Client,
) (*link.UC, error) {
	linkService, err := link.New(log, nil, store, authPermission, kratosClient, eventBus)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

func NewSitemapService(log logger.Logger, eventBus *bus.EventBus) (*sitemap.Service, error) {
	return sitemap.New(log, eventBus)
}

// TODO: refactoring. maybe drop this function
func NewRunRPCServer(runRPCServer *rpc.Server, _ *cqrs_rpc.LinkRPC, _ *link_rpc.LinkRPC) (*run.Response, error) {
	return run.Run(runRPCServer)
}

func NewLinkService(
	// Common
	log logger.Logger,
	config *config.Config,

	// Observability
	metrics *metrics.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	flightTrace *flight_trace.Recorder,

	// Security
	authPermission *authzed.Client,

	// Application
	linkService *link.UC,
	linkCQRSService *link_cqrs.Service,
	sitemapService *sitemap.Service,

	// Delivery
	run *run.Response,
	linkRPCServer *link_rpc.LinkRPC,
	linkCQRSRPCServer *cqrs_rpc.LinkRPC,
	sitemapRPCServer *sitemap_rpc.Sitemap,

	// Repository
	linkStore *crud.Store,

	// CQRS Repository
	cqsStore *cqs.Store,
	queryStore *query.Store,
) (*LinkService, error) {
	return &LinkService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Metrics:       metrics,
		PprofEndpoint: pprofHTTP,
		FlightTrace:   flightTrace,

		// Security
		authPermission: authPermission,

		// Application
		linkService:     linkService,
		linkCQRSService: linkCQRSService,
		sitemapService:  sitemapService,

		// Delivery
		run:               run,
		linkRPCServer:     linkRPCServer,
		linkCQRSRPCServer: linkCQRSRPCServer,
		sitemapRPCServer:  sitemapRPCServer,

		// Repository
		linkStore: linkStore,

		// CQRS
		cqsStore:   cqsStore,
		queryStore: queryStore,
	}, nil
}

func InitializeLinkService() (*LinkService, func(), error) {
	panic(wire.Build(LinkSet))
}

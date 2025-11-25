//go:generate go tool wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Link UC DI-package
*/
package link_di

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/authzed/authzed-go/v1"
	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus"
	shortctx "github.com/shortlink-org/go-sdk/context"
	"github.com/shortlink-org/go-sdk/flags"
	"github.com/shortlink-org/go-sdk/flight_trace"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/observability/tracing"
	"go.opentelemetry.io/otel/metric"
	api "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/cqrs/query"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud"
	cqrs "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/rpc/cqrs/link/v1"
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
	linkCQRSRPCServer *cqrs.LinkRPC
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

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(
	// Common
	DefaultSet,
	permission.New,
	NewPrometheusRegistry,
	NewMeterProvider,
	db.New,

	// Delivery
	NewWatermillMeterProvider,
	NewWatermillBackend,
	watermill.New,
	NewWatermillPublisher,
	rpc.InitServer,
	NewRPCClient,
	link_rpc.New,
	cqrs.New,
	sitemap_rpc.New,
	NewRunRPCServer,

	link_rpc.NewLinkServiceClient,
	// metadata_rpc.NewMetadataServiceClient,

	// Applications
	NewLinkApplication,
	link_cqrs.New,
	NewSitemapService,

	// repository
	crud.New,
	cqs.New,
	query.New,

	NewLinkService,
)

func NewPrometheusRegistry(metrics *metrics.Monitoring) *prometheus.Registry {
	return metrics.Prometheus
}

func NewMeterProvider(metrics *metrics.Monitoring) *api.MeterProvider {
	return metrics.Metrics
}

func NewWatermillMeterProvider(metrics *metrics.Monitoring) metric.MeterProvider {
	return metrics.Metrics
}

func NewWatermillBackend(ctx context.Context, log logger.Logger, cfg *config.Config) (watermill.Backend, error) {
	return watermill_kafka.New(ctx, log, cfg)
}

func NewRPCClient(
	ctx context.Context,
	log logger.Logger,
	cfg *config.Config,
	metrics *metrics.Monitoring,
	tracer trace.TracerProvider,
) (*grpc.ClientConn, func(), error) {
	// Initialize gRPC Client's interceptor.
	opts := []rpc.Option{
		rpc.WithSession(),
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

func NewWatermillPublisher(client *watermill.Client) message.Publisher {
	return client.Publisher
}

func NewLinkApplication(log logger.Logger, publisher message.Publisher, store *crud.Store, authPermission *authzed.Client) (*link.UC, error) {
	linkService, err := link.New(log, publisher, nil, store, authPermission)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

func NewSitemapService(log logger.Logger, publisher message.Publisher) (*sitemap.Service, error) {
	return sitemap.New(log, publisher)
}

// TODO: refactoring. maybe drop this function
func NewRunRPCServer(runRPCServer *rpc.Server, _ *cqrs.LinkRPC, _ *link_rpc.LinkRPC) (*run.Response, error) {
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
	linkCQRSRPCServer *cqrs.LinkRPC,
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

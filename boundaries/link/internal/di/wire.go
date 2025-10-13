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
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	api_mq "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/mq"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/cqrs/query"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud"
	cqrs "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/rpc/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/rpc/run"
	sitemap_rpc "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/rpc/sitemap/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/link"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/link_cqrs"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/sitemap"

	"github.com/shortlink-org/go-sdk/config"
	rpc "github.com/shortlink-org/go-sdk/grpc"
	"github.com/shortlink-org/go-sdk/logger"

	"github.com/shortlink-org/shortlink/pkg/di"
	mq_di "github.com/shortlink-org/shortlink/pkg/di/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/permission"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/store"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
)

type LinkService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Metrics       *metrics.Monitoring
	PprofEndpoint profiling.PprofEndpoint

	// Security
	authPermission *authzed.Client

	// Delivery
	linkMQ            *api_mq.Event
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

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(
	// Common
	di.DefaultSet,
	permission.New,
	store.New,
	NewPrometheusRegistry,

	// Delivery
	mq_di.New,
	api_mq.New,
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
	sitemap.New,

	// repository
	crud.New,
	cqs.New,
	query.New,

	NewLinkService,
)

func NewPrometheusRegistry(metrics *metrics.Monitoring) *prometheus.Registry {
	return metrics.Prometheus
}

func NewRPCClient(
	ctx context.Context,
	log logger.Logger,
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

	runRPCClient, cleanup, err := rpc.InitClient(ctx, log, opts...)
	if err != nil {
		return nil, nil, err
	}

	return runRPCClient, cleanup, nil
}

func NewLinkApplication(log logger.Logger, mq mq.MQ, store *crud.Store, authPermission *authzed.Client) (*link.UC, error) {
	linkService, err := link.New(log, mq, nil, store, authPermission)
	if err != nil {
		return nil, err
	}

	return linkService, nil
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

	// Security
	authPermission *authzed.Client,

	// Application
	linkService *link.UC,
	linkCQRSService *link_cqrs.Service,
	sitemapService *sitemap.Service,

	// Delivery
	linkMQ *api_mq.Event,
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
		linkMQ:            linkMQ,

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

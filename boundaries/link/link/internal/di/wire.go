//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Link UC DI-package
*/
package link_di

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

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
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	mq_di "github.com/shortlink-org/shortlink/pkg/di/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/store"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type LinkService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
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
	di.DefaultSet,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
	store.New,

	// Delivery
	api_mq.New,

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
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Observability
	monitoring *monitoring.Monitoring,
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
		Log:        log,
		Config:     config,
		AutoMaxPro: autoMaxProcsOption,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
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

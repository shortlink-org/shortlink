//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/application/link"
	"github.com/batazor/shortlink/internal/services/link/application/link_cqrs"
	"github.com/batazor/shortlink/internal/services/link/application/sitemap"
	api_domain "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	api_mq "github.com/batazor/shortlink/internal/services/link/infrastructure/mq"
	cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/run"
	sitemap_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/cqs"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/query"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	"github.com/batazor/shortlink/pkg/rpc"
)

type LinkService struct {
	Logger logger.Logger

	// Delivery
	linkMQ            *api_mq.Event
	run               *run.Response
	linkRPCServer     *link_rpc.Link
	linkCQRSRPCServer *cqrs.Link
	sitemapRPCServer  *sitemap_rpc.Sitemap

	// Application
	linkService     *link.Service
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
	// Delivery
	InitLinkMQ,

	NewLinkRPCServer,
	NewLinkCQRSRPCServer,
	NewSitemapRPCServer,
	NewRunRPCServer,

	NewLinkRPCClient,
	NewMetadataRPCClient,

	// applications
	NewLinkApplication,
	NewLinkCQRSApplication,
	NewSitemapApplication,

	// repository
	NewLinkStore,
	NewCQSLinkStore,
	NewQueryLinkStore,

	NewLinkService,
)

func InitLinkMQ(ctx context.Context, log logger.Logger, mq v1.MQ) (*api_mq.Event, error) {
	linkMQ, err := api_mq.New(mq, log)
	if err != nil {
		return nil, err
	}

	// Publish Event
	notify.Subscribe(api_domain.METHOD_ADD, linkMQ)

	return linkMQ, nil
}

func NewLinkStore(ctx context.Context, logger logger.Logger, db *db.Store) (*crud.Store, error) {
	store := &crud.Store{}
	linkStore, err := store.Use(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

func NewCQSLinkStore(ctx context.Context, logger logger.Logger, db *db.Store) (*cqs.Store, error) {
	store := &cqs.Store{}
	cqsStore, err := store.Use(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	return cqsStore, nil
}

func NewQueryLinkStore(ctx context.Context, logger logger.Logger, db *db.Store) (*query.Store, error) {
	store := &query.Store{}
	queryStore, err := store.Use(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	return queryStore, nil
}

func NewLinkApplication(logger logger.Logger, metadataService metadata_rpc.MetadataServiceClient, store *crud.Store) (*link.Service, error) {
	linkService, err := link.New(logger, metadataService, store)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

func NewLinkCQRSApplication(logger logger.Logger, cqsStore *cqs.Store, queryStore *query.Store) (*link_cqrs.Service, error) {
	linkCQRSService, err := link_cqrs.New(logger, cqsStore, queryStore)
	if err != nil {
		return nil, err
	}

	return linkCQRSService, nil
}

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (link_rpc.LinkServiceClient, error) {
	LinkServiceClient := link_rpc.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewSitemapApplication(logger logger.Logger, mq v1.MQ) (*sitemap.Service, error) {
	sitemapService, err := sitemap.New(logger, mq)
	if err != nil {
		return nil, err
	}

	return sitemapService, nil
}

func NewLinkCQRSRPCServer(runRPCServer *rpc.RPCServer, application *link_cqrs.Service, log logger.Logger) (*cqrs.Link, error) {
	linkRPCServer, err := cqrs.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewLinkRPCServer(runRPCServer *rpc.RPCServer, application *link.Service, log logger.Logger) (*link_rpc.Link, error) {
	linkRPCServer, err := link_rpc.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewSitemapRPCServer(runRPCServer *rpc.RPCServer, application *sitemap.Service, log logger.Logger) (*sitemap_rpc.Sitemap, error) {
	sitemapRPCServer, err := sitemap_rpc.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return sitemapRPCServer, nil
}

func NewRunRPCServer(runRPCServer *rpc.RPCServer, cqrsLinkRPC *cqrs.Link, linkRPC *link_rpc.Link) (*run.Response, error) {
	return run.Run(runRPCServer)
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (metadata_rpc.MetadataServiceClient, error) {
	metadataRPCClient := metadata_rpc.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewLinkService(
	log logger.Logger,

	// Application
	linkService *link.Service,
	linkCQRSService *link_cqrs.Service,
	sitemapService *sitemap.Service,

	// Delivery
	linkMQ *api_mq.Event,
	run *run.Response,
	linkRPCServer *link_rpc.Link,
	linkCQRSRPCServer *cqrs.Link,
	sitemapRPCServer *sitemap_rpc.Sitemap,

	// Repository
	linkStore *crud.Store,

	// CQRS Repository
	cqsStore *cqs.Store,
	queryStore *query.Store,
) (*LinkService, error) {
	return &LinkService{
		Logger: log,

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

func InitializeLinkService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq v1.MQ) (*LinkService, func(), error) {
	panic(wire.Build(LinkSet))
}

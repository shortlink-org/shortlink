//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/application/link"
	"github.com/batazor/shortlink/internal/services/link/application/link_cqrs"
	api_domain "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	api_mq "github.com/batazor/shortlink/internal/services/link/infrastructure/mq"
	cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	v12 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/run"
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
	linkRPCServer     *v12.Link
	linkCQRSRPCServer *cqrs.Link

	// Application
	linkService     *link.Service
	linkCQRSService *link_cqrs.Service

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
	NewRunRPCServer,
	NewMetadataRPCClient,

	// applications
	NewLinkApplication,
	NewLinkCQRSApplication,

	// repository
	NewLinkStore,
	NewCQSLinkStore,
	NewQueryLinkStore,

	NewLinkService,
)

func InitLinkMQ(ctx context.Context, log logger.Logger, mq mq.MQ) (*api_mq.Event, error) {
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

func NewLinkCQRSRPCServer(runRPCServer *rpc.RPCServer, application *link_cqrs.Service, log logger.Logger) (*cqrs.Link, error) {
	linkRPCServer, err := cqrs.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewLinkRPCServer(runRPCServer *rpc.RPCServer, application *link.Service, log logger.Logger) (*v12.Link, error) {
	linkRPCServer, err := v12.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewRunRPCServer(runRPCServer *rpc.RPCServer, cqrsLinkRPC *cqrs.Link, linkRPC *v12.Link) (*run.Response, error) {
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

	// Delivery
	linkMQ *api_mq.Event,
	run *run.Response,
	linkRPCServer *v12.Link,
	linkCQRSRPCServer *cqrs.Link,

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

		// Delivery
		run:               run,
		linkRPCServer:     linkRPCServer,
		linkCQRSRPCServer: linkCQRSRPCServer,
		linkMQ:            linkMQ,

		// Repository
		linkStore: linkStore,

		// CQRS
		cqsStore:   cqsStore,
		queryStore: queryStore,
	}, nil
}

func InitializeLinkService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq mq.MQ) (*LinkService, func(), error) {
	panic(wire.Build(LinkSet))
}

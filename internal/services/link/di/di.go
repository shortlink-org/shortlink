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
	api_domain "github.com/batazor/shortlink/internal/services/api/domain"
	"github.com/batazor/shortlink/internal/services/link/application/link"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/cqrs/cqs"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/cqrs/query"
	api_mq "github.com/batazor/shortlink/internal/services/link/infrastructure/mq"
	v12 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
	"github.com/batazor/shortlink/pkg/rpc"
)

type LinkService struct {
	Logger logger.Logger

	// Delivery
	linkRPCServer *v12.Link

	// Application
	service *link.Service

	// Repository
	linkMQ    *api_mq.Event
	linkStore *store.Store

	// CQRS
	cqsStore   *cqs.Store
	queryStore *query.Store
}

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(
	// infrastructure
	NewLinkRPCServer,
	NewMetadataRPCClient,
	InitLinkMQ,

	// applications
	NewLinkApplication,

	// repository
	NewLinkStore,
	NewCQSLinkStore,
	NewQueryLinkStore,

	NewLinkService,
)

func InitLinkMQ(ctx context.Context, log logger.Logger, mq mq.MQ) (*api_mq.Event, error) {
	linkMQ := &api_mq.Event{
		MQ: mq,
	}

	// Subscribe to Event
	notify.Subscribe(api_domain.METHOD_ADD, linkMQ)

	return linkMQ, nil
}

func NewLinkStore(ctx context.Context, logger logger.Logger, db *db.Store) (*store.Store, error) {
	store := &store.Store{}
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

func NewLinkApplication(logger logger.Logger, metadataService metadata_rpc.MetadataClient, store *store.Store, cqsStore *cqs.Store, queryStore *query.Store) (*link.Service, error) {
	linkService, err := link.New(logger, metadataService, store, cqsStore, queryStore)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

func NewLinkRPCServer(runRPCServer *rpc.RPCServer, application *link.Service, log logger.Logger) (*v12.Link, error) {
	linkRPCServer, err := v12.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (metadata_rpc.MetadataClient, error) {
	metadataRPCClient := metadata_rpc.NewMetadataClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewLinkService(
	log logger.Logger,
	linkRPCServer *v12.Link,
	service *link.Service,

	// Repository
	linkStore *store.Store,
	linkMQ *api_mq.Event,

	// CQRS
	cqsStore *cqs.Store,
	queryStore *query.Store,
) (*LinkService, error) {
	return &LinkService{
		Logger:        log,
		linkRPCServer: linkRPCServer,
		linkMQ:        linkMQ,
		service:       service,

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

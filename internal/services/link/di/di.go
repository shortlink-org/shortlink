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
	"github.com/batazor/shortlink/internal/services/link/application"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	api_mq "github.com/batazor/shortlink/internal/services/link/infrastructure/mq"
	v12 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	link_store "github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
	"github.com/batazor/shortlink/pkg/rpc"
)

type LinkService struct {
	Logger logger.Logger

	// Delivery
	linkRPCServer *v12.Link

	// Application
	service *link_application.Service

	// Repository
	linkStore *link_store.LinkStore
	linkMQ    *api_mq.Event
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

	NewLinkService,
)

func InitLinkMQ(ctx context.Context, log logger.Logger, mq mq.MQ) (*api_mq.Event, error) {
	linkMQ := &api_mq.Event{
		MQ: mq,
	}

	// Subscribe to Event
	notify.Subscribe(uint32(v1.LinkEvent_LINK_EVENT_ADD), linkMQ)

	return linkMQ, nil
}

func NewLinkStore(ctx context.Context, logger logger.Logger, db *db.Store) (*link_store.LinkStore, error) {
	store := &link_store.LinkStore{}
	linkStore, err := store.Use(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

func NewLinkApplication(logger logger.Logger, metadataService metadata_rpc.MetadataClient, store *link_store.LinkStore) (*link_application.Service, error) {
	linkService, err := link_application.New(logger, metadataService, store)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

func NewLinkRPCServer(runRPCServer *rpc.RPCServer, application *link_application.Service, log logger.Logger) (*v12.Link, error) {
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
	linkStore *link_store.LinkStore,
	service *link_application.Service,
	linkMQ *api_mq.Event,
) (*LinkService, error) {
	return &LinkService{
		Logger:        log,
		linkRPCServer: linkRPCServer,
		linkStore:     linkStore,
		linkMQ:        linkMQ,
		service:       service,
	}, nil
}

func InitializeLinkService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq mq.MQ) (*LinkService, func(), error) {
	panic(wire.Build(LinkSet))
}

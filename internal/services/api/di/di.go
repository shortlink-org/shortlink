//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/logger"
	api_application "github.com/batazor/shortlink/internal/services/api/application"
	v1 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	"github.com/batazor/shortlink/pkg/rpc"
)

type APIService struct {
	Logger logger.Logger

	// applications
	service *api_application.Server
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	// infrastructure
	NewLinkRPCClient,
	NewLinkCommandRPCClient,
	NewLinkQueryRPCClient,
	NewMetadataRPCClient,

	// applications
	NewAPIApplication,

	NewAPIService,
)

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (link_rpc.LinkServiceClient, error) {
	LinkServiceClient := link_rpc.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewLinkCommandRPCClient(runRPCClient *grpc.ClientConn) (v1.LinkCommandServiceClient, error) {
	LinkCommandRPCClient := v1.NewLinkCommandServiceClient(runRPCClient)
	return LinkCommandRPCClient, nil
}

func NewLinkQueryRPCClient(runRPCClient *grpc.ClientConn) (v1.LinkQueryServiceClient, error) {
	LinkQueryRPCClient := v1.NewLinkQueryServiceClient(runRPCClient)
	return LinkQueryRPCClient, nil
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (metadata_rpc.MetadataServiceClient, error) {
	metadataRPCClient := metadata_rpc.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewAPIApplication(
	ctx context.Context,
	logger logger.Logger,
	tracer *opentracing.Tracer,
	rpcServer *rpc.RPCServer,
	metadataClient metadata_rpc.MetadataServiceClient,
	linkServiceClient link_rpc.LinkServiceClient,
	linkCommandRPCClient v1.LinkCommandServiceClient,
	linkQueryRPCClient v1.LinkQueryServiceClient,
) (*api_application.Server, error) {
	// Run API server
	API := api_application.Server{
		MetadataClient:           metadataClient,
		LinkServiceClient:        linkServiceClient,
		LinkCommandServiceClient: linkCommandRPCClient,
		LinkQueryServiceClient:   linkQueryRPCClient,
	}

	apiService, err := API.RunAPIServer(ctx, logger, tracer, rpcServer)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewAPIService(
	log logger.Logger,

	service *api_application.Server,
) (*APIService, error) {
	return &APIService{
		Logger: log,

		service: service,
	}, nil
}

func InitializeAPIService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, tracer *opentracing.Tracer) (*APIService, func(), error) {
	panic(wire.Build(APISet))
}

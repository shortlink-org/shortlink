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
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
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
	NewLinkCommandRPCClient,
	NewLinkQueryRPCClient,
	NewMetadataRPCClient,

	// applications
	NewAPIApplication,

	NewAPIService,
)

func NewLinkCommandRPCClient(runRPCClient *grpc.ClientConn) (link_rpc.LinkCommandServiceClient, error) {
	LinkCommandRPCClient := link_rpc.NewLinkCommandServiceClient(runRPCClient)
	return LinkCommandRPCClient, nil
}

func NewLinkQueryRPCClient(runRPCClient *grpc.ClientConn) (link_rpc.LinkQueryServiceClient, error) {
	LinkQueryRPCClient := link_rpc.NewLinkQueryServiceClient(runRPCClient)
	return LinkQueryRPCClient, nil
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (metadata_rpc.MetadataClient, error) {
	metadataRPCClient := metadata_rpc.NewMetadataClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewAPIApplication(
	ctx context.Context,
	logger logger.Logger,
	tracer *opentracing.Tracer,
	rpcServer *rpc.RPCServer,
	metadataClient metadata_rpc.MetadataClient,
	linkCommandRPCClient link_rpc.LinkCommandServiceClient,
	linkQueryRPCClient link_rpc.LinkQueryServiceClient,
) (*api_application.Server, error) {
	// Run API server
	API := api_application.Server{
		MetadataClient:           metadataClient,
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

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/text/message"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/logger"
	api_application "github.com/batazor/shortlink/internal/services/api/application"
	link_cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	"github.com/batazor/shortlink/pkg/rpc"
)

type APIService struct {
	Logger logger.Logger

	// applications
	service *api_application.API
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	// infrastructure
	NewLinkRPCClient,
	NewLinkCommandRPCClient,
	NewLinkQueryRPCClient,
	NewSitemapServiceClient,
	NewMetadataRPCClient,

	// applications
	NewAPIApplication,

	NewAPIService,
)

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (link_rpc.LinkServiceClient, error) {
	LinkServiceClient := link_rpc.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewLinkCommandRPCClient(runRPCClient *grpc.ClientConn) (link_cqrs.LinkCommandServiceClient, error) {
	LinkCommandRPCClient := link_cqrs.NewLinkCommandServiceClient(runRPCClient)
	return LinkCommandRPCClient, nil
}

func NewLinkQueryRPCClient(runRPCClient *grpc.ClientConn) (link_cqrs.LinkQueryServiceClient, error) {
	LinkQueryRPCClient := link_cqrs.NewLinkQueryServiceClient(runRPCClient)
	return LinkQueryRPCClient, nil
}

func NewSitemapServiceClient(runRPCClient *grpc.ClientConn) (sitemap_rpc.SitemapServiceClient, error) {
	sitemapRPCClient := sitemap_rpc.NewSitemapServiceClient(runRPCClient)
	return sitemapRPCClient, nil
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (metadata_rpc.MetadataServiceClient, error) {
	metadataRPCClient := metadata_rpc.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewAPIApplication(
	ctx context.Context,
	i18n *message.Printer,
	logger logger.Logger,
	tracer *opentracing.Tracer,
	rpcServer *rpc.RPCServer,

	// delivery
	metadataClient metadata_rpc.MetadataServiceClient,
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) (*api_application.API, error) {
	// Run API server
	apiService, err := api_application.RunAPIServer(
		ctx,
		i18n,
		logger,
		tracer,
		rpcServer,

		// delivery
		link_rpc,
		link_command,
		link_query,
		sitemap_rpc,
	)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewAPIService(
	log logger.Logger,
	i18n *message.Printer,

	service *api_application.API,
) (*APIService, error) {
	return &APIService{
		Logger: log,

		service: service,
	}, nil
}

func InitializeAPIService(ctx context.Context, i18n *message.Printer, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, tracer *opentracing.Tracer) (*APIService, func(), error) {
	panic(wire.Build(APISet))
}

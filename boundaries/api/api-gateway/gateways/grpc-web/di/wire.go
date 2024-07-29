//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
API SERVICE DI-package
*/
package api_di

import (
	"context"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"

	"github.com/shortlink-org/shortlink/boundaries/api/api-gateway/gateways/grpc-web/infrastructure/server"
	api_application "github.com/shortlink-org/shortlink/boundaries/api/api-gateway/gateways/grpc-web/infrastructure/server/v1"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type APIService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Applications
	service *api_application.API

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	di.DefaultSet,

	// Delivery
	rpc.InitServer,
	rpc.InitClient,

	// Infrastructure
	NewLinkRPCClient,
	NewLinkCommandRPCClient,
	NewLinkQueryRPCClient,
	NewSitemapServiceClient,
	// NewMetadataRPCClient,

	// Applications
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

// func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (metadata_rpc.MetadataServiceClient, error) {
// 	metadataRPCClient := metadata_rpc.NewMetadataServiceClient(runRPCClient)
// 	return metadataRPCClient, nil
// }

func NewAPIApplication(
	// Common
	ctx context.Context,
	log logger.Logger,
	rpcServer *rpc.Server,
	tracer trace.TracerProvider,
	monitor *monitoring.Monitoring,

	// Delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) (*api_application.API, error) {
	// Run API server
	apiService, err := server.RunAPIServer(
		// Common
		ctx,
		log,
		rpcServer,
		tracer,
		monitor,

		// Delivery
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
	// Common
	log logger.Logger,
	config *config.Config,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	service *api_application.API,
) (*APIService, error) {
	return &APIService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		service: service,
	}, nil
}

func InitializeAPIService() (*APIService, func(), error) {
	panic(wire.Build(APISet))
}

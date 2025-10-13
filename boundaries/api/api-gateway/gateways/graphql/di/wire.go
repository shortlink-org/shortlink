//go:generate go tool wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
API SERVICE DI-package
*/
package api_di

import (
	"context"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/shortlink-org/go-sdk/config"
	rpc "github.com/shortlink-org/go-sdk/grpc"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/boundaries/api/api-gateway/gateways/graphql/infrastructure/server"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/permission"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
)

type APIService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Applications
	service *server.API

	// Observability
	Tracer        trace.TracerProvider
	Metrics       *metrics.Monitoring
	PprofEndpoint profiling.PprofEndpoint
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	di.DefaultSet,
	permission.New,

	// Delivery
	rpc.InitServer,
	NewRPCClient,

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

func NewRPCClient(
	ctx context.Context,
	log logger.Logger,
	metrics *metrics.Monitoring,
	tracer trace.TracerProvider,
) (*grpc.ClientConn, func(), error) {
	// Initialize gRPC Client's interceptor.
	opts := []rpc.Option{
		rpc.WithSession(),
		rpc.WithMetrics(metrics),
		rpc.WithTracer(tracer, metrics),
		rpc.WithTimeout(),
		rpc.WithLogger(log),
	}

	runRPCClient, cleanup, err := rpc.InitClient(ctx, log, opts...)
	if err != nil {
		return nil, nil, err
	}

	return runRPCClient, cleanup, nil
}

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
	monitor *metrics.Monitoring,

	// Delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) (*server.API, error) {
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

	// Observability
	metrics *metrics.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	service *server.API,
) (*APIService, error) {
	return &APIService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Metrics:       metrics,
		PprofEndpoint: pprofHTTP,

		service: service,
	}, nil
}

func InitializeAPIService() (*APIService, func(), error) {
	panic(wire.Build(APISet))
}

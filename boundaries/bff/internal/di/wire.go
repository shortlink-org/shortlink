//go:generate go tool wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
BFF Web Service DI-package
*/
package bff_web_di

import (
	"context"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"
	"google.golang.org/grpc"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"

	"github.com/shortlink-org/go-sdk/config"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/permission"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
	"github.com/shortlink-org/shortlink/pkg/rpc"

	api "github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/pkg/i18n"
)

type BFFWebService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	i18n       *message.Printer
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Delivery
	httpAPIServer *api.Server

	// Observability
	Tracer        trace.TracerProvider
	Metrics       *metrics.Monitoring
	PprofEndpoint profiling.PprofEndpoint
}

// BFFWebService =======================================================================================================
var BFFWebServiceSet = wire.NewSet(
	di.DefaultSet,
	permission.New,
	i18n.New,

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
	NewBFFWebService,
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
	return link_rpc.NewLinkServiceClient(runRPCClient), nil
}

func NewLinkCommandRPCClient(runRPCClient *grpc.ClientConn) (link_cqrs.LinkCommandServiceClient, error) {
	return link_cqrs.NewLinkCommandServiceClient(runRPCClient), nil
}

func NewLinkQueryRPCClient(runRPCClient *grpc.ClientConn) (link_cqrs.LinkQueryServiceClient, error) {
	return link_cqrs.NewLinkQueryServiceClient(runRPCClient), nil
}

func NewSitemapServiceClient(runRPCClient *grpc.ClientConn) (sitemap_rpc.SitemapServiceClient, error) {
	return sitemap_rpc.NewSitemapServiceClient(runRPCClient), nil
}

// func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (metadata_rpc.MetadataServiceClient, error) {
// 	return metadata_rpc.NewMetadataServiceClient(runRPCClient), nil
// }

func NewAPIApplication(
	// Common
	ctx context.Context,
	i18n *message.Printer,
	log logger.Logger,
	config *config.Config,
	autoMaxPro autoMaxPro.AutoMaxPro,

	// Observability
	tracer trace.TracerProvider,
	metrics *metrics.Monitoring,
	pprofEndpoint profiling.PprofEndpoint,

	// Infrastructure
	rpcServer *rpc.Server,
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) (*api.Server, error) {
	apiService, err := api.New(api.Config{
		// Common
		Ctx:    ctx,
		I18n:   i18n,
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Metrics:       metrics,
		PprofEndpoint: pprofEndpoint,
		AutoMaxPro:    autoMaxPro,

		// Delivery
		RpcServer: rpcServer,

		// Infrastructure
		Link_rpc:     link_rpc,
		Link_command: link_command,
		Link_query:   link_query,
		Sitemap_rpc:  sitemap_rpc,
	})
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewBFFWebService(
	// Common
	ctx context.Context,
	log logger.Logger,
	config *config.Config,
	autoMaxPro autoMaxPro.AutoMaxPro,

	// Observability
	tracer trace.TracerProvider,
	metrics *metrics.Monitoring,
	pprofEndpoint profiling.PprofEndpoint,

	// Delivery
	httpAPIServer *api.Server,
) *BFFWebService {
	return &BFFWebService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Metrics:       metrics,
		PprofEndpoint: pprofEndpoint,
		AutoMaxPro:    autoMaxPro,

		// Delivery
		httpAPIServer: httpAPIServer,
	}
}

func InitializeBFFWebService() (*BFFWebService, func(), error) {
	panic(wire.Build(BFFWebServiceSet))
}

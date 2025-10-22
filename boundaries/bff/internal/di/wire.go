//go:generate go tool wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
BFF Web Service DI-package
*/
package bff_di

import (
	"context"

	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus"
	shortctx "github.com/shortlink-org/go-sdk/context"
	"github.com/shortlink-org/go-sdk/flags"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/observability/tracing"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"
	"google.golang.org/grpc"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"

	"github.com/shortlink-org/go-sdk/config"
	rpc "github.com/shortlink-org/go-sdk/grpc"

	"github.com/shortlink-org/go-sdk/auth/permission"
	"github.com/shortlink-org/go-sdk/cache"
	"github.com/shortlink-org/go-sdk/observability/metrics"
	"github.com/shortlink-org/go-sdk/observability/profiling"

	api "github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/pkg/i18n"
)

type BFFWebService struct {
	// Common
	Log    logger.Logger
	Config *config.Config
	i18n   *message.Printer

	// Delivery
	httpAPIServer *api.Server

	// Observability
	Tracer        trace.TracerProvider
	Metrics       *metrics.Monitoring
	PprofEndpoint profiling.PprofEndpoint
}

// DefaultSet ==========================================================================================================
var DefaultSet = wire.NewSet(
	shortctx.New,
	flags.New,
	config.New,
	logger.NewDefault,
	tracing.New,
	metrics.New,
	cache.New,
	profiling.New,
)

// BFFWebService =======================================================================================================
var BFFWebServiceSet = wire.NewSet(
	DefaultSet,
	permission.New,
	i18n.New,
	NewPrometheusRegistry,
	NewMeterProvider,

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

func NewPrometheusRegistry(metrics *metrics.Monitoring) *prometheus.Registry {
	return metrics.Prometheus
}

func NewMeterProvider(metrics *metrics.Monitoring) *metric.MeterProvider {
	return metrics.Metrics
}

func NewRPCClient(
	ctx context.Context,
	log logger.Logger,
	metrics *metrics.Monitoring,
	tracer trace.TracerProvider,
) (*grpc.ClientConn, func(), error) {
	// Initialize gRPC Client's interceptor.
	opts := []rpc.Option{
		rpc.WithSession(),
		rpc.WithMetrics(metrics.Prometheus),
		rpc.WithTracer(tracer, metrics.Prometheus, metrics.Metrics),
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

		// Delivery
		RpcServer: rpcServer,

		// Infrastructure
		Link_rpc:     link_rpc,
		Link_command: link_command,
		Link_query:   link_query,
		Sitemap_rpc:  sitemap_rpc,
	}, log)
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

		// Delivery
		httpAPIServer: httpAPIServer,
	}
}

func InitializeBFFWebService() (*BFFWebService, func(), error) {
	panic(wire.Build(BFFWebServiceSet))
}

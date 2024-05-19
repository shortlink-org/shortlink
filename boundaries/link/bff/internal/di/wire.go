//go:generate wire
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

	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/pkg/i18n"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"

	api "github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http"
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
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro
}

// BFFWebService =======================================================================================================
var BFFWebServiceSet = wire.NewSet(
	di.DefaultSet,
	i18n.New,

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
	NewBFFWebService,
)

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
	monitoring *monitoring.Monitoring,
	pprofEndpoint profiling.PprofEndpoint,
	autoMaxPro autoMaxPro.AutoMaxPro,

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
		Monitoring:    monitoring,
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

	// Observability
	tracer trace.TracerProvider,
	monitoring *monitoring.Monitoring,
	pprofEndpoint profiling.PprofEndpoint,
	autoMaxPro autoMaxPro.AutoMaxPro,

	// Delivery
	httpAPIServer *api.Server,
) *BFFWebService {
	return &BFFWebService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofEndpoint,
		AutoMaxPro:    autoMaxPro,

		// Delivery
		httpAPIServer: httpAPIServer,
	}
}

func InitializeBFFWebService() (*BFFWebService, func(), error) {
	panic(wire.Build(BFFWebServiceSet))
}

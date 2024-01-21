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

	"github.com/shortlink-org/shortlink/internal/boundaries/api/bff-web/internal/i18n"
	link_cqrs "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/sitemap/v1"
	metadata_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/infrastructure/rpc/metadata/v1"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/rpc"

	api "github.com/shortlink-org/shortlink/internal/boundaries/api/bff-web/infrastructure/http"
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
	NewMetadataRPCClient,

	// Applications
	NewAPIApplication,
	NewBFFWebService,
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

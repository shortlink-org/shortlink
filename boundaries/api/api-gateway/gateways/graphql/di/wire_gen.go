// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api_di

import (
	"context"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/boundaries/api/api-gateway/gateways/graphql/infrastructure/server"
	v1_2 "github.com/shortlink-org/shortlink/boundaries/link/link/infrastructure/rpc/cqrs/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/infrastructure/rpc/link/v1"
	v1_3 "github.com/shortlink-org/shortlink/boundaries/link/link/infrastructure/rpc/sitemap/v1"
	v1_4 "github.com/shortlink-org/shortlink/boundaries/link/metadata/infrastructure/rpc/metadata/v1"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitializeAPIService() (*APIService, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracerProvider, cleanup3, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	monitoringMonitoring, cleanup4, err := monitoring.New(context, logger, tracerProvider)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint, err := profiling.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup5, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, err := rpc.InitServer(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup6, err := rpc.InitClient(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkServiceClient, err := NewLinkRPCClient(clientConn)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkCommandServiceClient, err := NewLinkCommandRPCClient(clientConn)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkQueryServiceClient, err := NewLinkQueryRPCClient(clientConn)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sitemapServiceClient, err := NewSitemapServiceClient(clientConn)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	api, err := NewAPIApplication(context, logger, server, tracerProvider, monitoringMonitoring, linkServiceClient, linkCommandServiceClient, linkQueryServiceClient, sitemapServiceClient)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	apiService, err := NewAPIService(logger, configConfig, monitoringMonitoring, tracerProvider, pprofEndpoint, autoMaxProAutoMaxPro, api)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return apiService, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type APIService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Applications
	service *server.API

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(di.DefaultSet, rpc.InitServer, rpc.InitClient, NewLinkRPCClient,
	NewLinkCommandRPCClient,
	NewLinkQueryRPCClient,
	NewSitemapServiceClient,
	NewMetadataRPCClient,

	NewAPIApplication,
	NewAPIService,
)

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (v1.LinkServiceClient, error) {
	LinkServiceClient := v1.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewLinkCommandRPCClient(runRPCClient *grpc.ClientConn) (v1_2.LinkCommandServiceClient, error) {
	LinkCommandRPCClient := v1_2.NewLinkCommandServiceClient(runRPCClient)
	return LinkCommandRPCClient, nil
}

func NewLinkQueryRPCClient(runRPCClient *grpc.ClientConn) (v1_2.LinkQueryServiceClient, error) {
	LinkQueryRPCClient := v1_2.NewLinkQueryServiceClient(runRPCClient)
	return LinkQueryRPCClient, nil
}

func NewSitemapServiceClient(runRPCClient *grpc.ClientConn) (v1_3.SitemapServiceClient, error) {
	sitemapRPCClient := v1_3.NewSitemapServiceClient(runRPCClient)
	return sitemapRPCClient, nil
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (v1_4.MetadataServiceClient, error) {
	metadataRPCClient := v1_4.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewAPIApplication(ctx2 context.Context,

	log logger.Logger,
	rpcServer *rpc.Server,
	tracer trace.TracerProvider,
	monitor *monitoring.Monitoring,

	link_rpc v1.LinkServiceClient,
	link_command v1_2.LinkCommandServiceClient,
	link_query v1_2.LinkQueryServiceClient,
	sitemap_rpc v1_3.SitemapServiceClient,
) (*server.API, error) {

	apiService, err := server.RunAPIServer(ctx2, log,
		rpcServer,
		tracer,
		monitor,

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

	log logger.Logger, config2 *config.Config, monitoring2 *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	service *server.API,
) (*APIService, error) {
	return &APIService{

		Log:    log,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		service: service,
	}, nil
}

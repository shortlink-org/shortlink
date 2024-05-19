//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
WS SERVICE DI-package
*/
package ws_di

import (
	"context"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/boundaries/api/api-gateway/gateways/ws/infrustracture/ws"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
)

type WSService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Applications
	service *ws.WS

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro
}

// WSService ===========================================================================================================
var WSSet = wire.NewSet(
	di.DefaultSet,

	// Infrastructure
	NewWSServer,

	NewWSService,
)

func NewWSServer(ctx context.Context, log logger.Logger) (*ws.WS, error) {
	wsServer := &ws.WS{}
	return wsServer.Run(ctx, log)
}

func NewWSService(
	// Common
	log logger.Logger,
	config *config.Config,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	service *ws.WS,
) (*WSService, error) {
	return &WSService{
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

func InitializeWSService() (*WSService, func(), error) {
	panic(wire.Build(WSSet))
}

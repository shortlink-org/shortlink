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

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/services/api-gateway/gateways/ws/infrustracture/ws"
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
		Logger: log,
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

//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
OMS Cart Worker DI-package
*/
package oms_cart_worker_di

import (
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/cart/cart_worker"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/temporal"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
)

type OMSCartWorkerService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint

	// Temporal
	temporalClient client.Client
	cartWorker     worker.Worker
}

// OMSCartWorkerService ================================================================================================
var OMSCartWorkerSet = wire.NewSet(
	di.DefaultSet,

	// Temporal
	temporal.New,
	cart_worker.New,

	NewOMSCartWorkerService,
)

func NewOMSCartWorkerService(
	// Common
	log logger.Logger,
	config *config.Config,
	autoMaxPro autoMaxPro.AutoMaxPro,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	// Temporal
	temporalClient client.Client,
	cartWorker worker.Worker,
) (*OMSCartWorkerService, error) {
	return &OMSCartWorkerService{
		// Common
		Log:        log,
		Config:     config,
		AutoMaxPro: autoMaxPro,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,

		// Temporal
		temporalClient: temporalClient,
		cartWorker:     cartWorker,
	}, nil
}

func InitializeOMSCartWorkerService() (*OMSCartWorkerService, func(), error) {
	panic(wire.Build(OMSCartWorkerSet))
}

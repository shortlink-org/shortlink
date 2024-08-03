//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
OMS DI-package
*/
package oms_di

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"go.temporal.io/sdk/client"

	cartRPC "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/run"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/temporal"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type OMSService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint

	// Security
	authPermission *authzed.Client

	// Delivery
	run           *run.Response
	cartRPCServer *cartRPC.CartRPC

	// Temporal
	temporalClient client.Client
}

// OMSService ==========================================================================================================
var OMSSet = wire.NewSet(
	di.DefaultSet,
	rpc.InitServer,

	// Delivery
	cartRPC.New,
	NewRunRPCServer,

	NewOMSService,

	// Temporal
	temporal.New,
)

// TODO: refactoring. maybe drop this function
func NewRunRPCServer(runRPCServer *rpc.Server, _ *cartRPC.CartRPC) (*run.Response, error) {
	return run.Run(runRPCServer)
}

func NewOMSService(
	// Common
	log logger.Logger,
	config *config.Config,
	autoMaxPro autoMaxPro.AutoMaxPro,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	// Security
	authPermission *authzed.Client,

	// Delivery
	run *run.Response,
	cartRPCServer *cartRPC.CartRPC,

	// Temporal
	temporalClient client.Client,
) (*OMSService, error) {
	return &OMSService{
		// Common
		Log:        log,
		Config:     config,
		AutoMaxPro: autoMaxPro,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,

		// Security
		// TODO: enable later
		// authPermission: authPermission,

		// Delivery
		run:           run,
		cartRPCServer: cartRPCServer,

		// Temporal
		temporalClient: temporalClient,
	}, nil
}

func InitializeOMSService() (*OMSService, func(), error) {
	panic(wire.Build(OMSSet))
}

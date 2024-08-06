// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package oms_order_worker_di

import (
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/order/order_worker"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/context"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/temporal"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"go.opentelemetry.io/otel/trace"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// Injectors from wire.go:

func InitializeOMSOrderWorkerService() (*OMSOrderWorkerService, func(), error) {
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
	autoMaxProAutoMaxPro, cleanup3, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracerProvider, cleanup4, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	monitoringMonitoring, cleanup5, err := monitoring.New(context, logger, tracerProvider)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint, err := profiling.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	client, err := temporal.New(logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	worker, err := order_worker.New(context, client, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	omsOrderWorkerService, err := NewOMSOrderWorkerService(logger, configConfig, autoMaxProAutoMaxPro, monitoringMonitoring, tracerProvider, pprofEndpoint, client, worker)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return omsOrderWorkerService, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type OMSOrderWorkerService struct {
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
	orderWorker    worker.Worker
}

// OMSOrderWorkerService ================================================================================================
var OMSOrderWorkerSet = wire.NewSet(di.DefaultSet, temporal.New, order_worker.New, NewOMSOrderWorkerService)

func NewOMSOrderWorkerService(

	log logger.Logger, config2 *config.Config, autoMaxPro2 autoMaxPro.AutoMaxPro, monitoring2 *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	temporalClient client.Client,
	OrderWorker worker.Worker,
) (*OMSOrderWorkerService, error) {
	return &OMSOrderWorkerService{

		Log:        log,
		Config:     config2,
		AutoMaxPro: autoMaxPro2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,

		temporalClient: temporalClient,
		orderWorker:    OrderWorker,
	}, nil
}
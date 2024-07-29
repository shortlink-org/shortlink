//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package csi_di

import (
	"context"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
)

// Service - heplers
type Service struct {
	// Common
	Ctx        context.Context
	Log        logger.Logger
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
}

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

// CSI =================================================================================================================
var CSISet = wire.NewSet(di.DefaultSet, NewSCIDriver)

func NewSCIDriver(
	// Common
	log logger.Logger,
	ctx context.Context,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Observability
	monitor *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
) (*Service, error) {
	return &Service{
		// Common
		Ctx: ctx,
		Log: log,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitor,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,
	}, nil
}

func InitializeSCIDriver() (*Service, func(), error) {
	panic(wire.Build(CSISet))
}

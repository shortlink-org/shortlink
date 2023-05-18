//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
BFF Web Service DI-package
*/
package bff_web_di

import (
	"net/http"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

type BillingService struct {
	// Common
	Logger logger.Logger
	Config *config.Config

	// Observability
	Tracer        *trace.TracerProvider
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro
}

// BFFWebService =======================================================================================================
var BFFWebServiceSet = wire.NewSet(
	di.DefaultSet,

	NewBFFWebServiceSet,
)

func NewBFFWebServiceSet(
	// Common
	logger logger.Logger,
	config *config.Config,

	// Observability
	tracer *trace.TracerProvider,
	monitoring *http.ServeMux,
	pprofEndpoint profiling.PprofEndpoint,
	autoMaxPro autoMaxPro.AutoMaxPro,
) *BillingService {
	return &BillingService{
		// Common
		Logger: logger,
		Config: config,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofEndpoint,
		AutoMaxPro:    autoMaxPro,
	}
}

func InitializeBFFWebService() (*BillingService, func(), error) {
	panic(wire.Build(BFFWebServiceSet))
}

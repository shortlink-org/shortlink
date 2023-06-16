//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
BFF Web Service DI-package
*/
package bff_web_di

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"

	api "github.com/shortlink-org/shortlink/internal/services/bff-web/infrastructure/http"
)

type BFFWebService struct {
	// Common
	Logger logger.Logger
	Config *config.Config

	// Observability
	Tracer        *trace.TracerProvider
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	httpAPIServer *api.Server
}

// BFFWebService =======================================================================================================
var BFFWebServiceSet = wire.NewSet(
	di.DefaultSet,

	// Infrastructure
	BFFWebAPIService,

	NewBFFWebService,
)

func BFFWebAPIService(
	// Common
	ctx context.Context,
	logger logger.Logger,
	tracer *trace.TracerProvider,
) (*api.Server, error) {
	// Run API server
	API := api.Server{}
	apiService, err := API.Run(ctx, logger, tracer)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewBFFWebService(
	// Common
	ctx context.Context,
	logger logger.Logger,
	config *config.Config,

	// Observability
	tracer *trace.TracerProvider,
	monitoring *http.ServeMux,
	pprofEndpoint profiling.PprofEndpoint,
	autoMaxPro autoMaxPro.AutoMaxPro,

	// Delivery
	httpAPIServer *api.Server,
) *BFFWebService {
	return &BFFWebService{
		// Common
		Logger: logger,
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

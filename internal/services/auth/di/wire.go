//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Auth Service DI-package
*/
package auth_di

import (
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/pkg/auth"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/services/auth/services/permission"
)

type LinkService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Jobs
	authPermission *auth.Auth
}

// AuthService =========================================================================================================
var AuthSet = wire.NewSet(
	di.DefaultSet,

	// Jobs
	permission.Permission,

	NewAuthService,
)

func NewAuthService(
	// Common
	log logger.Logger,
	config *config.Config,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Jobs
	authPermission *auth.Auth,
) (*LinkService, error) {
	return &LinkService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		// Jobs
		authPermission: authPermission,
	}, nil
}

func InitializeAuthService() (*LinkService, func(), error) {
	panic(wire.Build(AuthSet))
}

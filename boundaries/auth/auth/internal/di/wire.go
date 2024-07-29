//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Auth Service DI-package
*/
package auth_di

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/boundaries/auth/auth/internal/services/permission"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
)

type AuthService struct {
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

	// Application
	permissionService *permission.Service
}

// AuthService =========================================================================================================
var AuthSet = wire.NewSet(
	di.DefaultSet,

	// Application
	permission.New,

	NewAuthService,
)

func NewAuthService(
	// Common
	log logger.Logger,
	config *config.Config,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	// Security
	authPermission *authzed.Client,

	// Application
	permissionService *permission.Service,
) (*AuthService, error) {
	return &AuthService{
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

		// Application
		permissionService: permissionService,
	}, nil
}

func InitializeAuthService() (*AuthService, func(), error) {
	panic(wire.Build(AuthSet))
}

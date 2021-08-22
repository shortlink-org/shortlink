//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/services/logger/di"
)

type ServiceLogger struct {
	Service

	loggerService *di.LoggerService
}

// InitLoggerService ===================================================================================================
func InitLoggerService(ctx context.Context, log logger.Logger, mq v1.MQ) (*di.LoggerService, func(), error) {
	return di.InitializeLoggerService(ctx, log, mq)
}

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(
	DefaultSet,
	sentry.New,
	monitoring.New,
	mq_di.New,
	InitLoggerService,
	NewLoggerService,
)

func NewLoggerService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	mq v1.MQ,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	loggerService *di.LoggerService,
) (*ServiceLogger, error) {
	return &ServiceLogger{
		Service: Service{
			Ctx:        ctx,
			Log:        log,
			MQ:         mq,
			Tracer:     tracer,
			Monitoring: monitoring,
		},
		loggerService: loggerService,
	}, nil
}

func InitializeLoggerService() (*ServiceLogger, func(), error) {
	panic(wire.Build(LoggerSet))
}

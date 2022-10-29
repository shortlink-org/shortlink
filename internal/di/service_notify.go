//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/batazor/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/pkg/config"
	"github.com/batazor/shortlink/internal/di/pkg/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/pkg/mq"
	"github.com/batazor/shortlink/internal/di/pkg/sentry"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
)

// NotifyService ==========================================================================================================
var NotifySet = wire.NewSet(
	DefaultSet,
	mq_di.New,
	sentry.New,
	monitoring.New,
	NewNotifyService,
)

func NewNotifyService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	mq v1.MQ,
	monitoring *http.ServeMux,
	tracer *trace.TracerProvider,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		Ctx:        ctx,
		Log:        log,
		MQ:         mq,
		Tracer:     tracer,
		Monitoring: monitoring,
	}, nil
}

func InitializeNotifyService() (*Service, func(), error) {
	panic(wire.Build(NotifySet))
}

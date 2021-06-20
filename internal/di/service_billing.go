//+build wireinject
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
	"github.com/batazor/shortlink/internal/pkg/mq"
)

// BillingService =======================================================================================================
var BillingSet = wire.NewSet(
	DefaultSet,
	mq_di.New,
	sentry.New,
	monitoring.New,
	NewBillingService,
)

func NewBillingService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	mq mq.MQ,
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

func InitializeBillingService() (*Service, func(), error) {
	panic(wire.Build(BillingSet))
}

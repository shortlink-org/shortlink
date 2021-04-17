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

// BotService ==========================================================================================================
var BotSet = wire.NewSet(
	DefaultSet,
	mq_di.New,
	sentry.New,
	monitoring.New,
	NewBotService,
)

func NewBotService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	mq mq.MQ,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
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

func InitializeBotService() (*Service, func(), error) {
	panic(wire.Build(BotSet))
}

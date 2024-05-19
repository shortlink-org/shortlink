//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package notify_di

import (
	"context"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"

	"github.com/shortlink-org/shortlink/boundaries/notification/notify/internal/application"
	"github.com/shortlink-org/shortlink/boundaries/notification/notify/internal/infrastructure/slack"
	"github.com/shortlink-org/shortlink/boundaries/notification/notify/internal/infrastructure/smtp"
	"github.com/shortlink-org/shortlink/boundaries/notification/notify/internal/infrastructure/telegram"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	mq_di "github.com/shortlink-org/shortlink/pkg/di/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

// Service - heplers
type Service struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Bot
	slack    *slack.Bot
	telegram *telegram.Bot
	smtp     *smtp.Bot

	// Application
	botService *application.Bot
}

// InitSlack - Init slack bot
func InitSlack(ctx context.Context) *slack.Bot {
	slackBot := &slack.Bot{}
	if err := slackBot.Init(); err != nil {
		return nil
	}

	return slackBot
}

// InitTelegram - Init telegram bot
func InitTelegram(ctx context.Context) *telegram.Bot {
	telegramBot := &telegram.Bot{}
	if err := telegramBot.Init(); err != nil {
		return nil
	}

	return telegramBot
}

// InitSMTP - Init SMTP bot
func InitSMTP(ctx context.Context) *smtp.Bot {
	smtpBot := &smtp.Bot{}
	if err := smtpBot.Init(); err != nil {
		return nil
	}

	return smtpBot
}

// FullBotService ======================================================================================================
var NotifySet = wire.NewSet(
	di.DefaultSet,
	mq_di.New,

	InitSlack,
	InitTelegram,
	InitSMTP,

	// Applications
	NewBotApplication,

	NewBotService,
)

func NewBotApplication(ctx context.Context, log logger.Logger, dataBus mq.MQ) (*application.Bot, error) {
	bot, err := application.New(dataBus, log)
	if err != nil {
		return nil, err
	}
	bot.Use(ctx)

	return bot, nil
}

func NewBotService(
	// Common
	log logger.Logger,
	config *config.Config,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Bots
	slack *slack.Bot,
	telegram *telegram.Bot,
	smtp *smtp.Bot,

	// Application
	bot *application.Bot,
) (*Service, error) {
	return &Service{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		// Bots
		slack:    slack,
		telegram: telegram,
		smtp:     smtp,

		// Application
		botService: bot,
	}, nil
}

func InitializeFullBotService() (*Service, func(), error) {
	panic(wire.Build(NotifySet))
}

//go:generate wire
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package notify_di

import (
	"context"

	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/di"
	mq_di "github.com/batazor/shortlink/internal/di/pkg/mq"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/services/notify/application"
	"github.com/batazor/shortlink/internal/services/notify/infrastructure/slack"
	"github.com/batazor/shortlink/internal/services/notify/infrastructure/smtp"
	"github.com/batazor/shortlink/internal/services/notify/infrastructure/telegram"
)

// Service - heplers
type Service struct {
	// Common
	Log logger.Logger

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

func NewBotApplication(ctx context.Context, logger logger.Logger, mq v1.MQ) (*application.Bot, error) {
	bot, err := application.New(mq, logger)
	if err != nil {
		return nil, err
	}
	bot.Use(ctx)

	return bot, nil
}

func NewBotService(
	// Common
	log logger.Logger,

	slack *slack.Bot,
	telegram *telegram.Bot,
	smtp *smtp.Bot,

	bot *application.Bot,
) (*Service, error) {
	return &Service{
		// Common
		Log: log,

		slack:      slack,
		telegram:   telegram,
		smtp:       smtp,
		botService: bot,
	}, nil
}

func InitializeFullBotService() (*Service, func(), error) {
	panic(wire.Build(NotifySet))
}

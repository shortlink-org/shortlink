//go:generate wire
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/services/notify/infrastructure/slack"
	"github.com/batazor/shortlink/internal/services/notify/infrastructure/smtp"
	"github.com/batazor/shortlink/internal/services/notify/infrastructure/telegram"
)

// Service - heplers
type Service struct {
	slack    *slack.Bot
	telegram *telegram.Bot
	smtp     *smtp.Bot
}

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
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
var FullBotSet = wire.NewSet(NewContext, InitSlack, InitTelegram, InitSMTP, NewBotService)

func NewBotService(slack *slack.Bot, telegram *telegram.Bot, smtp *smtp.Bot) (*Service, error) {
	return &Service{
		slack,
		telegram,
		smtp,
	}, nil
}

func InitializeFullBotService() (*Service, func(), error) {
	panic(wire.Build(FullBotSet))
}

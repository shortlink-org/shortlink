//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
)

// BotService ==========================================================================================================
var BotSet = wire.NewSet(DefaultSet, NewBotService, InitMQ)

func NewBotService(log logger.Logger, mq mq.MQ, autoMaxProcsOption diAutoMaxPro) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq,
	}, nil
}

func InitializeBotService() (*Service, func(), error) {
	panic(wire.Build(BotSet))
}

//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
)

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(DefaultSet, NewLoggerService, InitMQ)

func NewLoggerService(log logger.Logger, mq mq.MQ, autoMaxProcsOption diAutoMaxPro) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq,
	}, nil
}

func InitializeLoggerService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(LoggerSet))
}

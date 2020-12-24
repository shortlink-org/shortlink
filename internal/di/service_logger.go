//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
)

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(
	DefaultSet,
	mq_di.New,
	monitoring.New,
	NewLoggerService,
)

func NewLoggerService(log logger.Logger, mq mq.MQ, autoMaxProcsOption autoMaxPro.AutoMaxPro) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq,
	}, nil
}

func InitializeLoggerService() (*Service, func(), error) {
	panic(wire.Build(LoggerSet))
}

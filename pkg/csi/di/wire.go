//go:generate wire
//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
)

// Service - heplers
type Service struct {
	Log logger.Logger
}

func InitLogger(ctx context.Context) (logger.Logger, func(), error) {
	viper.SetDefault("LOG_LEVEL", logger.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := logger.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	log, err := logger.NewLogger(logger.Zap, conf)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		// flushes buffer, if any
		_ = log.Close() // nolint errcheck
	}

	return log, cleanup, nil
}

// CSI =================================================================================================================
var FullBotSet = wire.NewSet(InitLogger, NewSCIDriver)

func NewSCIDriver(log logger.Logger) (*Service, error) {
	return &Service{
		Log: log,
	}, nil
}

func InitializeSCIDriver(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(FullBotSet))
}

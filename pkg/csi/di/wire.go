//go:generate wire
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/logger"
)

// Service - heplers
type Service struct {
	Ctx context.Context
	Log logger.Logger
}

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

// Logger ==============================================================================================================
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
var FullBotSet = wire.NewSet(NewContext, InitLogger, NewSCIDriver)

func NewSCIDriver(log logger.Logger, ctx context.Context) (*Service, error) {
	return &Service{
		Ctx: ctx,
		Log: log,
	}, nil
}

func InitializeSCIDriver() (*Service, func(), error) {
	panic(wire.Build(FullBotSet))
}

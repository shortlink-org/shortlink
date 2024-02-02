package logger_di

import (
	"context"
	"time"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/config"
)

func New(_ context.Context) (logger.Logger, func(), error) {
	viper.SetDefault("LOG_LEVEL", config.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := config.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	log, err := logger.New(logger.Zap, conf)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		// flushes buffer, if any
		_ = log.Close() //nolint:errcheck // ignore
	}

	return log, cleanup, nil
}

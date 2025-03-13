package logger_di

import (
	"context"
	"time"

	"github.com/spf13/viper"

	error_di "github.com/shortlink-org/shortlink/pkg/di/pkg/error"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/config"
)

// New creates a new logger instance
//
//nolint:ireturn // It's made by design
func New(_ context.Context) (logger.Logger, func(), error) {
	viper.SetDefault("LOG_LEVEL", config.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := config.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	log, err := logger.New(logger.Zap, conf)
	if err != nil {
		return nil, nil, &error_di.BaseError{Err: err}
	}

	cleanup := func() {
		// flushes buffer, if any
		_ = log.Close() //nolint:errcheck // ignore
	}

	return log, cleanup, nil
}

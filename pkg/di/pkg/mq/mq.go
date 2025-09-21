package mq_di

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/logger"
	error_di "github.com/shortlink-org/shortlink/pkg/di/pkg/error"
	"github.com/shortlink-org/shortlink/pkg/mq"
)

// New creates a new MQ instance
//
//nolint:ireturn // It's made by design
func New(ctx context.Context, log logger.Logger) (mq.MQ, error) {
	viper.SetDefault("MQ_ENABLED", "false") // Enabled MQ

	if !viper.GetBool("MQ_ENABLED") {
		//nolint:nilnil // It's made by design
		return nil, nil
	}

	var service mq.DataBus

	dataBus, err := service.Use(ctx, log)
	if err != nil {
		return nil, &error_di.BaseError{Err: err}
	}

	return dataBus, nil
}

package mq_di

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
)

func New(ctx context.Context, log logger.Logger) (mq.MQ, error) {
	viper.SetDefault("MQ_ENABLED", "false") // Enabled MQ

	if viper.GetBool("MQ_ENABLED") {
		var service mq.DataBus
		dataBus, err := service.Use(ctx, log)
		if err != nil {
			return nil, err
		}

		return dataBus, nil
	}

	log.Warn("MQ disabled")

	return nil, nil
}

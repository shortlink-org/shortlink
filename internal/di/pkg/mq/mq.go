package mq_di

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
)

func New(ctx context.Context, log logger.Logger) (*mq.DataBus, func(), error) {
	viper.SetDefault("MQ_ENABLED", "false") // Enabled MQ

	if viper.GetBool("MQ_ENABLED") {
		var service mq.DataBus
		dataBus, err := service.Use(ctx, log)
		if err != nil {
			return nil, func() {}, err
		}

		cleanup := func() {
			if err := dataBus.Close(); err != nil {
				log.Error(err.Error())
			}
		}

		return dataBus, cleanup, nil
	}

	log.Warn("MQ disabled")

	return nil, func() {}, nil
}

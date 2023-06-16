package mq_di

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/internal/pkg/mq"
)

func New(ctx context.Context, log logger.Logger) (*v1.DataBus, func(), error) {
	viper.SetDefault("MQ_ENABLED", "false") // Enabled MQ

	if viper.GetBool("MQ_ENABLED") {
		var service v1.DataBus
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
	} else {
		log.Warn("MQ disabled")
	}

	return nil, func() {}, nil
}

//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
)

// Store ===============================================================================================================
// InitStore return db
func InitStore(ctx context.Context, log logger.Logger) (*db.Store, func(), error) {
	var st db.Store
	db, err := st.Use(ctx, log)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Store.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db, cleanup, nil
}

// MQ ==================================================================================================================
func InitMQ(ctx context.Context, log logger.Logger) (mq.MQ, func(), error) {
	viper.SetDefault("MQ_ENABLED", "false") // Enabled MQ-service

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

	return nil, func() {}, nil
}

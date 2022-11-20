/*
Message Queue
*/
package v1

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/kafka"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/nats"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/rabbit"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/redis"
)

// Use return implementation of MQ
func (mq *DataBus) Use(ctx context.Context, log logger.Logger) (MQ, error) {
	// Set configuration
	mq.setConfig()

	switch mq.typeMQ {
	case "kafka":
		mq.mq = &kafka.Kafka{}
	case "nats":
		mq.mq = nats.New()
	case "rabbitmq":
		mq.mq = rabbit.New(log)
	case "redis":
		mq.mq = redis.New()
	default:
		mq.mq = &kafka.Kafka{}
	}

	if err := mq.mq.Init(ctx); err != nil {
		return nil, err
	}

	log.Info("run MQ", field.Fields{
		"mq": mq.typeMQ,
	})

	return mq.mq, nil
}

// setConfig - set configuration
func (mq *DataBus) setConfig() {
	viper.SetDefault("MQ_TYPE", "rabbitmq") // Select: kafka, rabbitmq, nats, redis
	mq.typeMQ = viper.GetString("MQ_TYPE")
}

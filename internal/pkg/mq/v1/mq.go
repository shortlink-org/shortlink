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
)

// Use return implementation of MQ
func (mq *DataBus) Use(ctx context.Context, log logger.Logger) (MQ, error) {
	// Set configuration
	mq.setConfig()

	switch mq.typeMQ {
	case "kafka":
		mq.mq = &kafka.Kafka{}
	case "nats":
		mq.mq = &nats.NATS{}
	case "rabbitmq":
		mq.mq = &rabbit.RabbitMQ{
			Log: log,
		}
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
func (mq *DataBus) setConfig() { // nolint unused
	viper.SetDefault("MQ_TYPE", "rabbitmq") // Select: kafka, rabbitmq, nats
	mq.typeMQ = viper.GetString("MQ_TYPE")
}

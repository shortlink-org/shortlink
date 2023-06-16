/*
Message Queue
*/
package mq

import (
	"context"

	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/kafka"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/nats"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/rabbit"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/redis"
)

// Use return implementation of MQ
func (mq *DataBus) Use(ctx context.Context, log logger.Logger) (*DataBus, error) {
	// Set configuration
	mq.setConfig()

	// Set logger
	mq.log = log

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

	return mq, nil
}

// Subscribe - subscribe to a topic
func (mq *DataBus) Subscribe(ctx context.Context, target string, message query.Response) error {
	mq.log.Info("subscribe to topic", field.Fields{
		"topic": target,
	})

	return mq.mq.Subscribe(ctx, target, message)
}

// UnSubscribe - unsubscribe to a topic
func (mq *DataBus) UnSubscribe(target string) error {
	mq.log.Info("unsubscribe to topic", field.Fields{
		"topic": target,
	})

	return mq.mq.UnSubscribe(target)
}

// Publish - publish to a topic
func (mq *DataBus) Publish(ctx context.Context, target string, key []byte, payload []byte) error {
	mq.log.Info("publish to topic", field.Fields{
		"topic": target,
	})

	return mq.mq.Publish(ctx, target, key, payload)
}

// Close - close connection
func (mq *DataBus) Close() error {
	mq.log.Info("close MQ", field.Fields{
		"mq": mq.typeMQ,
	})

	return mq.mq.Close()
}

// setConfig - set configuration
func (mq *DataBus) setConfig() {
	viper.SetDefault("MQ_TYPE", "rabbitmq") // Select: kafka, rabbitmq, nats, redis
	mq.typeMQ = viper.GetString("MQ_TYPE")
}

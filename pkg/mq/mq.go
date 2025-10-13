/*
Message Queue
*/
package mq

import (
	"context"
	"log/slog"

	"github.com/spf13/viper"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/pkg/mq/kafka"
	"github.com/shortlink-org/shortlink/pkg/mq/nats"
	"github.com/shortlink-org/shortlink/pkg/mq/query"
	"github.com/shortlink-org/shortlink/pkg/mq/rabbit"
	"github.com/shortlink-org/shortlink/pkg/mq/redis"
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

	if err := mq.Init(ctx, log); err != nil {
		return nil, err
	}

	return mq, nil
}

// Init - init connection
func (mq *DataBus) Init(ctx context.Context, log logger.Logger) error {
	err := mq.mq.Init(ctx, log)
	if err != nil {
		return err
	}

	mq.log.Info("run MQ", slog.String("mq", mq.typeMQ))

	return nil
}

// Subscribe - subscribe to a topic
func (mq *DataBus) Subscribe(ctx context.Context, target string, message query.Response) error {
	mq.log.Info("subscribe to topic",
		slog.String("topic", target),
	)

	return mq.mq.Subscribe(ctx, target, message)
}

// UnSubscribe - unsubscribe to a topic
func (mq *DataBus) UnSubscribe(target string) error {
	mq.log.Info("unsubscribe to topic",
		slog.String("topic", target),
	)

	return mq.mq.UnSubscribe(target)
}

// Publish - publish to a topic
func (mq *DataBus) Publish(ctx context.Context, target string, key, payload []byte) error {
	mq.log.Info("publish to topic",
		slog.String("topic", target),
	)

	return mq.mq.Publish(ctx, target, key, payload)
}

// setConfig - set configuration
func (mq *DataBus) setConfig() {
	viper.SetDefault("MQ_TYPE", "rabbitmq") // Select: kafka, rabbitmq, nats, redis
	mq.typeMQ = viper.GetString("MQ_TYPE")
}

/*
Message Queue
*/

package mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq/kafka"
	"github.com/batazor/shortlink/internal/mq/nats"
	"github.com/batazor/shortlink/internal/mq/query"
	"github.com/batazor/shortlink/internal/notify"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

// Use return implementation of MQ
func (mq *DataBus) Use(ctx context.Context, log logger.Logger) MQ { // nolint unused
	// Set configuration
	mq.setConfig()

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, mq)

	switch mq.typeMQ {
	case "kafka":
		mq.mq = &kafka.Kafka{}
	case "nats":
		mq.mq = &nats.NATS{}
	default:
		mq.mq = &kafka.Kafka{}
	}

	if err := mq.mq.Init(ctx); err != nil {
		panic(err)
	}

	log.Info("run MQ", logger.Fields{
		"mq": mq.typeMQ,
	})

	return mq.mq
}

// setConfig - set configuration
func (mq *DataBus) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("MQ_TYPE", "kafka")
	mq.typeMQ = viper.GetString("MQ_TYPE")
}

// Notify ...
func (mq *DataBus) Notify(event int, payload interface{}) *notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		// TODO: send []byte
		msg := payload.(link.Link) // nolint errcheck
		data, err := proto.Marshal(&msg)
		if err != nil {
			return &notify.Response{
				Name:    "RESPONSE_MQ_ADD",
				Payload: nil,
				Error:   err,
			}
		}

		err = mq.mq.Publish(query.Message{
			Key:     nil,
			Payload: data,
		})
		return &notify.Response{
			Name:    "RESPONSE_MQ_ADD",
			Payload: nil,
			Error:   err,
		}
	case api_type.METHOD_GET:
		panic("implement me")
	case api_type.METHOD_LIST:
		panic("implement me")
	case api_type.METHOD_UPDATE:
		panic("implement me")
	case api_type.METHOD_DELETE:
		panic("implement me")
	}

	return nil
}

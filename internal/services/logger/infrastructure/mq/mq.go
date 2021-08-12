/*
MQ Endpoint
*/

package logger_mq

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/batazor/shortlink/internal/pkg/logger"
	v12 "github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/logger/application"
)

type Event struct {
	mq  v12.MQ
	log logger.Logger

	service *logger_application.Service
}

func New(mq v12.MQ, log logger.Logger) (*Event, error) {
	amqpConfig := amqp.NewDurableQueueConfig(viper.GetString("MQ_RABBIT_URI"))
	amqpConfig.Exchange = amqp.ExchangeConfig{
		GenerateName: func(topic string) string {
			return topic
		},
		Type:        "fanout",
		Durable:     true,
		AutoDeleted: false,
		Internal:    false,
		NoWait:      false,
		Arguments:   nil,
	}

	subscriber, err := amqp.NewSubscriber(
		// This config is based on this example: https://www.rabbitmq.com/tutorials/tutorial-two-go.html
		// It works as a simple queue.
		//
		// If you want to implement a Pub/Sub style service instead, check
		// https://watermill.io/pubsubs/amqp/#amqp-consumer-groups
		amqpConfig,
		watermill.NewStdLogger(true, true),
	)
	if err != nil {
		return nil, err
	}

	messages, err := subscriber.Subscribe(context.Background(), "shortlink.link.event")
	if err != nil {
		panic(err)
	}

	go process(messages)

	event := &Event{
		mq:  mq,
		log: log,
	}

	// Subscribe
	//event.Subscribe()

	return event, nil
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}

func (e *Event) Subscribe() {
	getEventNewLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		if err := e.mq.Subscribe("shortlink.link.event", getEventNewLink); err != nil {
			e.log.Error(err.Error())
		}
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan

			// Convert: []byte to link.Link
			myLink := &v1.Link{}
			if err := proto.Unmarshal(msg.Body, myLink); err != nil {
				e.log.Error(fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				msg.Context.Done()
				continue
			}

			e.service.Log(msg.Context, myLink)
			msg.Context.Done()
		}
	}()
}

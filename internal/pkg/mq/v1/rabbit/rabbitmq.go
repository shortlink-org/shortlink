package rabbit

import (
	"context"
	"fmt"
	"sync"

	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
)

type RabbitMQ struct {
	wg sync.Mutex

	URI  string
	conn *Connection
	ch   *Channel

	reconnectTime int

	log logger.Logger
}

func New(log logger.Logger) *RabbitMQ {
	return &RabbitMQ{
		log: log,
	}
}

func (mq *RabbitMQ) Init(_ context.Context) error {
	// Set configuration
	mq.setConfig()

	// connect to RabbitMQ server
	err := mq.Dial()
	if err != nil {
		return err
	}

	// create a channel
	mq.ch, err = mq.conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

func (mq *RabbitMQ) Publish(ctx context.Context, target string, message query.Message) error {
	mq.wg.Lock()
	defer mq.wg.Unlock()

	sp := opentracing.SpanFromContext(ctx)
	defer sp.Finish()

	// create exchange
	err := mq.ch.ExchangeDeclare(
		target,   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        message.Payload,
		Headers:     make(amqp.Table),
	}

	// Inject the span context into the AMQP header.
	err = opentracing.GlobalTracer().Inject(sp.Context(), opentracing.TextMap, amqpHeadersCarrier(msg.Headers))
	if err != nil {
		mq.log.Warn(err.Error())
	}

	err = mq.ch.Publish(
		target, // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		msg,
	)
	if err != nil {
		return err
	}

	return nil
}

func (mq *RabbitMQ) Subscribe(target string, message query.Response) error {
	// create a exchange
	err := mq.ch.ExchangeDeclare(
		target,   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	// create a queue
	q, err := mq.ch.QueueDeclare(
		fmt.Sprintf("%s-%s", target, viper.GetString("SERVICE_NAME")), // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = mq.ch.QueueBind(
		q.Name,
		"*",
		target,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := mq.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	g := errgroup.Group{}

	g.Go(func() error {
		for msg := range msgs {
			ctx := context.Background()

			// Extract the span context out of the AMQP header.
			spanCtx, errSpan := opentracing.GlobalTracer().Extract(opentracing.TextMap, amqpHeadersCarrier(msg.Headers))
			if errSpan != nil {
				mq.log.Warn(errSpan.Error())
			}

			span := opentracing.StartSpan(
				"AMQP: ConsumeMessage",
				opentracing.FollowsFrom(spanCtx),
			)
			span.SetTag("queue", q.Name)

			// Update the context with the span for the subsequent reference.
			ctx = opentracing.ContextWithSpan(ctx, span)

			message.Chan <- query.ResponseMessage{
				Body:    msg.Body,
				Context: ctx,
			}

			span.Finish()
		}

		return nil
	})

	return g.Wait()
}

func (mq *RabbitMQ) UnSubscribe(target string) error {
	mq.wg.Lock()
	defer mq.wg.Unlock()

	return nil
}

func (mq *RabbitMQ) Close() error {
	if err := mq.conn.Close(); err != nil {
		return err
	}

	if err := mq.ch.Close(); err != nil {
		return err
	}

	return nil
}

// setConfig - Construct a new RabbitMQ configuration.
func (mq *RabbitMQ) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_RABBIT_URI", "amqp://localhost:5672") // RabbitMQ URI
	viper.SetDefault("MQ_RECONNECT_DELAY_SECONDS", 3)          // RabbitMQ reconnects after delay seconds

	mq.URI = viper.GetString("MQ_RABBIT_URI")
	mq.reconnectTime = viper.GetInt("MQ_RECONNECT_DELAY_SECONDS")
}

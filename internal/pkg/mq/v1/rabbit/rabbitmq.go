package rabbit

import (
	"context"
	"fmt"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"golang.org/x/sync/errgroup"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/v1/query"
)

type RabbitMQ struct {
	log           logger.Logger
	conn          *Connection
	ch            *Channel
	URI           string
	reconnectTime int
	wg            sync.Mutex
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

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        message.Payload,
		Headers:     make(amqp.Table),
	}

	// Inject the span context into the AMQP header.
	tc := propagation.TraceContext{}
	tc.Inject(ctx, amqpHeadersCarrier(msg.Headers))

	err := mq.ch.PublishWithContext(
		ctx,
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
			tc := propagation.TraceContext{}
			spanCtx := tc.Extract(ctx, amqpHeadersCarrier(msg.Headers))

			spanCtx, span := otel.Tracer("AMQP").Start(spanCtx, "ConsumeMessage")
			span.SetAttributes(attribute.String("queue", q.Name))

			message.Chan <- query.ResponseMessage{
				Body:    msg.Body,
				Context: spanCtx,
			}

			span.End()
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
	// RabbitMQ reconnects after delay seconds
	viper.SetDefault("MQ_RECONNECT_DELAY_SECONDS", 3) // nolint:gomnd

	mq.URI = viper.GetString("MQ_RABBIT_URI")
	mq.reconnectTime = viper.GetInt("MQ_RECONNECT_DELAY_SECONDS")
}

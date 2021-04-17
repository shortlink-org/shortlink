package rabbit

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/query"
)

type RabbitMQ struct {
	URI  string
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue

	Log logger.Logger
}

func (mq *RabbitMQ) Init(_ context.Context) error {
	var err error

	// Set configuration
	mq.setConfig()

	// connect to RabbitMQ server
	mq.conn, err = amqp.Dial(mq.URI)
	if err != nil {
		return err
	}

	// create a channel
	mq.ch, err = mq.conn.Channel()
	if err != nil {
		return err
	}

	// create a exchange
	err = mq.ch.ExchangeDeclare(
		"shortlink", // name
		"fanout",    // type
		false,       // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	// create a queue
	mq.q, err = mq.ch.QueueDeclare(
		"shortlink", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	err = mq.ch.QueueBind(
		mq.q.Name,
		"*",
		"shortlink",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (mq *RabbitMQ) Publish(ctx context.Context, message query.Message) error {
	sp := opentracing.SpanFromContext(ctx)
	defer sp.Finish()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        message.Payload,
		Headers:     make(amqp.Table),
	}

	// Inject the span context into the AMQP header.
	err := opentracing.GlobalTracer().Inject(sp.Context(), opentracing.TextMap, amqpHeadersCarrier(msg.Headers))
	if err != nil {
		mq.Log.Warn(err.Error())
	}

	err = mq.ch.Publish(
		string(message.Key), // exchange
		mq.q.Name,           // routing key
		false,               // mandatory
		false,               // immediate
		msg,
	)
	if err != nil {
		return err
	}

	return nil
}

func (mq *RabbitMQ) Subscribe(message query.Response) error {
	msgs, err := mq.ch.Consume(
		mq.q.Name, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			ctx := context.Background()

			// Extract the span context out of the AMQP header.
			spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, amqpHeadersCarrier(msg.Headers))
			if err != nil {
				mq.Log.Warn(err.Error())
			}

			span := opentracing.StartSpan(
				"AMQP: ConsumeMessage",
				opentracing.FollowsFrom(spanCtx),
			)
			span.SetTag("queue", mq.q.Name)

			// Update the context with the span for the subsequent reference.
			ctx = opentracing.ContextWithSpan(ctx, span)

			message.Chan <- query.ResponseMessage{
				Body:    msg.Body,
				Context: ctx,
			}

			span.Finish()
		}
	}()

	return nil
}

func (mq *RabbitMQ) UnSubscribe() error {
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

	mq.URI = viper.GetString("MQ_RABBIT_URI")
}

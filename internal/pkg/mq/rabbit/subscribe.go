package rabbit

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"golang.org/x/sync/errgroup"

	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

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
	mq.mu.Lock()
	defer mq.mu.Unlock()

	return nil
}

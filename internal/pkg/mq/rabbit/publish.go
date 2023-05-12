package rabbit

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
	"go.opentelemetry.io/otel/propagation"
)

func (mq *RabbitMQ) Publish(ctx context.Context, target string, message query.Message) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

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

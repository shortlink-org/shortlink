package rabbit

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/propagation"
)

func (mq *MQ) Publish(ctx context.Context, target string, routingKey, payload []byte) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        payload,
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

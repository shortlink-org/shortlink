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

func (mq *MQ) Subscribe(ctx context.Context, target string, message query.Response) error {
	queueName := fmt.Sprintf("%s-%s", target, viper.GetString("SERVICE_NAME"))

	q, err := mq.ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	err = mq.ch.QueueBind(q.Name, "*", target, false, nil)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	msgs, err := mq.ch.Consume(ctx, q.Name, "", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %w", err)
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					return nil
				}

				spanCtx := propagation.TraceContext{}.Extract(ctx, amqpHeadersCarrier(msg.Headers))
				spanCtx, span := otel.Tracer("AMQP").Start(spanCtx, "ConsumeMessage")
				span.SetAttributes(attribute.String("queue", q.Name))

				message.Chan <- query.ResponseMessage{
					Body:    msg.Body,
					Context: spanCtx,
				}

				span.End()
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	return g.Wait()
}

func (mq *MQ) UnSubscribe(target string) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	err := mq.ch.QueueUnbind(fmt.Sprintf("%s-%s", target, viper.GetString("SERVICE_NAME")), "*", target, nil)
	if err != nil {
		return fmt.Errorf("failed to unbind queue: %w", err)
	}

	return nil
}

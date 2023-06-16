package kafka

import (
	"github.com/Shopify/sarama"

	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	// response channel
	ch query.Response
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited.
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		session.MarkMessage(message, "")

		consumer.ch.Chan <- query.ResponseMessage{
			Body: message.Value,
		}
	}

	return nil
}

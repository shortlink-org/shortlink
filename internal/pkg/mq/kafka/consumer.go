package kafka

import (
	"github.com/IBM/sarama"

	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	// response channel
	ch query.Response

	session *sarama.ConsumerGroupSession
	ready   chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(s sarama.ConsumerGroupSession) error {
	consumer.session = &s

	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			{
				if !ok {
					return nil
				}

				session.MarkMessage(message, "")

				consumer.ch.Chan <- query.ResponseMessage{
					Body: message.Value,
				}
			}
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}

	return nil
}

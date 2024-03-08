package transactional_outbox

import (
	"context"

	v1 "github.com/shortlink-org/shortlink/pkg/pattern/transactional_outbox/domain/transactional_outbox/v1"
)

type TransactionalOutbox interface {
	// Publish publishes a message to the outbox
	Publish(ctx context.Context, msg v1.OutboxMessage) error

	// Poll polls the outbox for messages
	Poll(ctx context.Context) ([]v1.Outbox, error)
}

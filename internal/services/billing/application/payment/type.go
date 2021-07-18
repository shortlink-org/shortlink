package payment_application

import (
	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

type Payment struct {
	eventsourcing.AggregateHandler
	billing.Payment
}

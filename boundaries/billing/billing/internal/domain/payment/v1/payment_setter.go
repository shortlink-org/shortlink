package v1

import (
	"context"

	"github.com/segmentio/encoding/json"

	eventsourcing "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing/domain/eventsourcing/v1"
)

// ApplyEventPaymentCreated applies the EventPaymentCreated event
func (p *Payment) ApplyEventPaymentCreated(ctx context.Context, event *eventsourcing.Event) error {
	var payload EventPaymentCreated
	if err := json.Unmarshal([]byte(event.GetPayload()), &payload); err != nil {
		return err
	}

	p.id = payload.Id
	p.name = payload.Name
	p.status = payload.Status
	p.userId = payload.UserId
	return nil
}

// ApplyEventPaymentApproved applies the EventPaymentApproved event
func (p *Payment) ApplyEventPaymentApproved(ctx context.Context, event *eventsourcing.Event) error {
	var payload EventPaymentApproved
	if err := json.Unmarshal([]byte(event.GetPayload()), &payload); err != nil {
		return err
	}

	p.status = payload.Status
	return nil
}

// ApplyEventPaymentClosed applies the EventPaymentClosed event
func (p *Payment) ApplyEventPaymentClosed(ctx context.Context, event *eventsourcing.Event) error {
	var payload EventPaymentClosed
	if err := json.Unmarshal([]byte(event.GetPayload()), &payload); err != nil {
		return err
	}

	p.status = payload.Status
	return nil
}

// ApplyEventPaymentRejected applies the EventPaymentRejected event
func (p *Payment) ApplyEventPaymentRejected(ctx context.Context, event *eventsourcing.Event) error {
	var payload EventPaymentRejected
	if err := json.Unmarshal([]byte(event.GetPayload()), &payload); err != nil {
		return err
	}

	p.status = payload.Status
	return nil
}

// ApplyEventBalanceUpdated applies the EventBalanceUpdated event
func (p *Payment) ApplyEventBalanceUpdated(ctx context.Context, event *eventsourcing.Event) error {
	if p.status != StatusPayment_STATUS_PAYMENT_APPROVE {
		return &IncorrectStatusOfPaymentError{Status: p.status.String()}
	}

	var payload EventBalanceUpdated
	if err := json.Unmarshal([]byte(event.GetPayload()), &payload); err != nil {
		return err
	}

	p.amount += payload.Amount

	return nil
}

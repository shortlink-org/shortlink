package payment_application

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/payment/v1"
	eventsourcing "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing/domain/eventsourcing/v1"
)

// ApplyChange to payment
func (p *Payment) ApplyChange(ctx context.Context, event *eventsourcing.Event) error {
	switch event.GetType() {
	case billing.Event_EVENT_PAYMENT_CREATED.String():
		return p.Payment.ApplyEventPaymentCreated(ctx, event)
	case billing.Event_EVENT_PAYMENT_APPROVED.String():
		return p.Payment.ApplyEventPaymentApproved(ctx, event)
	case billing.Event_EVENT_PAYMENT_CLOSED.String():
		return p.Payment.ApplyEventPaymentClosed(ctx, event)
	case billing.Event_EVENT_PAYMENT_REJECTED.String():
		return p.Payment.ApplyEventPaymentRejected(ctx, event)
	case billing.Event_EVENT_BALANCE_UPDATED.String():
		return p.Payment.ApplyEventBalanceUpdated(ctx, event)
	default:
		return &NotFoundEventError{Type: event.GetType()}
	}
}

// HandleCommand create events and validate based on such a command
func (p *Payment) HandleCommand(ctx context.Context, command *eventsourcing.BaseCommand) error {
	event := &eventsourcing.Event{
		AggregateId:   p.Payment.GetId().String(),
		AggregateType: "Payment",
	}

	// start tracing
	_, span := otel.Tracer("event sourcing").Start(ctx, "HandleCommand")
	span.SetAttributes(attribute.String("aggregate_id", p.Payment.GetId().String()))
	defer span.End()

	switch t := command.GetType(); {
	case t == billing.Command_COMMAND_PAYMENT_CREATE.String():
		event.AggregateId = command.GetAggregateId()
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_PAYMENT_CREATED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_PAYMENT_CREATED.String()))
	case t == billing.Command_COMMAND_PAYMENT_APPROVE.String():
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_PAYMENT_APPROVED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_PAYMENT_APPROVED.String()))
	case t == billing.Command_COMMAND_PAYMENT_CLOSE.String():
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_PAYMENT_CLOSED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_PAYMENT_CLOSED.String()))
	case t == billing.Command_COMMAND_PAYMENT_REJECT.String():
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_PAYMENT_REJECTED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_PAYMENT_REJECTED.String()))
	case t == billing.Command_COMMAND_BALANCE_UPDATE.String():
		event.Payload = command.GetPayload()
		event.Type = billing.Event_EVENT_BALANCE_UPDATED.String()

		span.SetAttributes(attribute.String("event_type", billing.Event_EVENT_BALANCE_UPDATED.String()))
	default:
		return &NotFoundCommandError{Type: t}
	}

	err := p.ApplyChangeHelper(ctx, p, event, true)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

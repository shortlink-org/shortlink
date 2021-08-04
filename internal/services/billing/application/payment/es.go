package payment_application

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/protobuf/encoding/protojson"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

// ApplyChange to payment
//gocyclo:ignore
func (p *Payment) ApplyChange(event *eventsourcing.Event) error {
	switch t := event.Type; {
	case t == billing.Event_EVENT_PAYMENT_CREATED.String():
		var payload billing.EventPaymentCreated
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return err
		}

		p.Payment.Id = payload.Id
		p.Name = payload.Name
		p.Status = payload.Status
		p.UserId = payload.UserId
	case t == billing.Event_EVENT_PAYMENT_APPROVED.String():
		var payload billing.EventPaymentApproved
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return err
		}

		p.Status = payload.Status
	case t == billing.Event_EVENT_PAYMENT_CLOSED.String():
		var payload billing.EventPaymentClosed
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return err
		}

		p.Status = payload.Status
	case t == billing.Event_EVENT_PAYMENT_REJECTED.String():
		var payload billing.EventPaymentRejected
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return err
		}

		p.Status = payload.Status
	case t == billing.Event_EVENT_BALANCE_UPDATED.String():
		var payload billing.EventBalanceUpdated
		err := protojson.Unmarshal([]byte(event.Payload), &payload)
		if err != nil {
			return err
		}

		p.Amount += payload.Amount
	default:
		return fmt.Errorf("Not found event with type: %s", event.Type)
	}

	return nil
}

// HandleCommand create events and validate based on such command
func (p *Payment) HandleCommand(ctx context.Context, command *eventsourcing.BaseCommand) error {
	event := &eventsourcing.Event{
		AggregateId:   p.Payment.Id,
		AggregateType: "Payment",
	}

	// start tracing
	span, _ := opentracing.StartSpanFromContext(ctx, "event: HandleCommand")
	span.SetTag("aggregate id", p.Payment.Id)
	defer span.Finish()

	switch t := command.GetType(); {
	case t == billing.Command_COMMAND_PAYMENT_CREATE.String():
		event.AggregateId = command.AggregateId
		event.Payload = command.Payload
		event.Type = billing.Event_EVENT_PAYMENT_CREATED.String()

		span.SetTag("event type", billing.Event_EVENT_PAYMENT_CREATED.String())
	case t == billing.Command_COMMAND_PAYMENT_APPROVE.String():
		event.Payload = command.Payload

		span.SetTag("event type", billing.Event_EVENT_PAYMENT_APPROVED.String())
	case t == billing.Command_COMMAND_PAYMENT_CLOSE.String():
		event.Payload = command.Payload
		event.Type = billing.Event_EVENT_PAYMENT_CLOSED.String()

		span.SetTag("event type", billing.Event_EVENT_PAYMENT_CLOSED.String())
	case t == billing.Command_COMMAND_PAYMENT_REJECTE.String():
		event.Payload = command.Payload

		span.SetTag("event type", billing.Event_EVENT_PAYMENT_REJECTED.String())
	case t == billing.Command_COMMAND_BALANCE_UPDATE.String():
		event.Payload = command.Payload
		event.Type = billing.Event_EVENT_BALANCE_UPDATED.String()

		span.SetTag("event type", billing.Event_EVENT_BALANCE_UPDATED.String())
	}

	err := p.ApplyChangeHelper(p, event, true)
	if err != nil {
		span.SetTag("error", true)
		span.SetTag("message", err.Error())
		return err
	}

	return nil
}

package payment_application

import (
	"context"

	"github.com/google/uuid"
	"github.com/segmentio/encoding/json"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/payment/v1"
	eventsourcing "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing/domain/eventsourcing/v1"
)

func CommandPaymentCreate(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	aggregateId := uuid.New()

	billing.NewPaymentBuilder().
		SetId(aggregateId).
		SetAmount(in.GetAmount()).
		SetName(in.GetName()).
		SetUserId(in.GetUserId()).
		SetStatus(billing.StatusPayment_STATUS_PAYMENT_NEW)

	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "UpdateBalance")
	span.SetAttributes(attribute.String("aggregate id", aggregateId.String()))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_PAYMENT_CREATE.String()))
	defer span.End()

	payload, err := json.Marshal(in)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	span.SetAttributes(attribute.String("log", string(payload)))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_CREATE.String(),
		AggregateId:   aggregateId.String(),
		AggregateType: "Payment",
		Version:       0,
		Payload:       string(payload),
	}, nil
}

func CommandPaymentUpdateBalance(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "UpdateBalance")
	span.SetAttributes(attribute.String("aggregate id", in.GetId().String()))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_BALANCE_UPDATE.String()))
	defer span.End()

	payload, err := json.Marshal(in)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	span.SetAttributes(attribute.String("log", string(payload)))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_BALANCE_UPDATE.String(),
		AggregateId:   in.GetId().String(),
		AggregateType: "Payment",
		Version:       1,
		Payload:       string(payload),
	}, nil
}

func CommandPaymentClose(ctx context.Context, id uuid.UUID) (*eventsourcing.BaseCommand, error) {
	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "PaymentClose")
	span.SetAttributes(attribute.String("aggregate id", id.String()))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_PAYMENT_CLOSE.String()))
	defer span.End()

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_CLOSE.String(),
		AggregateId:   id.String(),
		AggregateType: "Payment",
		Version:       1,
		Payload:       id.String(),
	}, nil
}

func CommandPaymentApprove(ctx context.Context, id uuid.UUID) (*eventsourcing.BaseCommand, error) {
	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "PaymentApprove")
	span.SetAttributes(attribute.String("aggregate id", id.String()))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_PAYMENT_APPROVE.String()))
	defer span.End()

	span.SetAttributes(attribute.String("aggregate id", id.String()))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_APPROVE.String(),
		AggregateId:   id.String(),
		AggregateType: "Payment",
		Version:       1,
		Payload:       id.String(),
	}, nil
}

func CommandPaymentReject(ctx context.Context, id uuid.UUID) (*eventsourcing.BaseCommand, error) {
	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "PaymentReject")
	span.SetAttributes(attribute.String("aggregate id", id.String()))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_PAYMENT_REJECT.String()))
	defer span.End()

	span.SetAttributes(attribute.String("aggregate id", id.String()))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_REJECT.String(),
		AggregateId:   id.String(),
		AggregateType: "Payment",
		Version:       1,
		Payload:       id.String(),
	}, nil
}

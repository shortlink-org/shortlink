package payment_application

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/protobuf/encoding/protojson"

	eventsourcing "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/v1"
	billing "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/payment/v1"
)

func CommandPaymentCreate(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	aggregateId := uuid.New().String()
	in.Status = billing.StatusPayment_STATUS_PAYMENT_NEW
	in.Id = aggregateId

	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "UpdateBalance")
	span.SetAttributes(attribute.String("aggregate id", aggregateId))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_PAYMENT_CREATE.String()))
	defer span.End()

	payload, err := protojson.Marshal(in)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	span.SetAttributes(attribute.String("log", string(payload)))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_CREATE.String(),
		AggregateId:   aggregateId,
		AggregateType: "Payment",
		Version:       0,
		Payload:       string(payload),
	}, nil
}

func CommandPaymentUpdateBalance(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "UpdateBalance")
	span.SetAttributes(attribute.String("aggregate id", in.Id))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_BALANCE_UPDATE.String()))
	defer span.End()

	payload, err := protojson.Marshal(in)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	span.SetAttributes(attribute.String("log", string(payload)))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_BALANCE_UPDATE.String(),
		AggregateId:   in.Id,
		AggregateType: "Payment",
		Version:       1,
		Payload:       string(payload),
	}, nil
}

func CommandPaymentClose(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	in.Status = billing.StatusPayment_STATUS_PAYMENT_CLOSE

	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "PaymentClose")
	span.SetAttributes(attribute.String("aggregate id", in.Id))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_PAYMENT_CLOSE.String()))
	defer span.End()

	payload, err := protojson.Marshal(in)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	span.SetAttributes(attribute.String("log", string(payload)))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_CLOSE.String(),
		AggregateId:   in.Id,
		AggregateType: "Payment",
		Version:       1,
		Payload:       string(payload),
	}, nil
}

func CommandPaymentApprove(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	in.Status = billing.StatusPayment_STATUS_PAYMENT_APPROVE

	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "PaymentApprove")
	span.SetAttributes(attribute.String("aggregate id", in.Id))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_PAYMENT_APPROVE.String()))
	defer span.End()

	payload, err := protojson.Marshal(in)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	span.SetAttributes(attribute.String("log", string(payload)))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_APPROVE.String(),
		AggregateId:   in.Id,
		AggregateType: "Payment",
		Version:       1,
		Payload:       string(payload),
	}, nil
}

func CommandPaymentReject(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	in.Status = billing.StatusPayment_STATUS_PAYMENT_REJECT

	// start tracing
	_, span := otel.Tracer("command").Start(ctx, "PaymentReject")
	span.SetAttributes(attribute.String("aggregate id", in.Id))
	span.SetAttributes(attribute.String("command type", billing.Command_COMMAND_PAYMENT_REJECT.String()))
	defer span.End()

	payload, err := protojson.Marshal(in)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	span.SetAttributes(attribute.String("log", string(payload)))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_REJECT.String(),
		AggregateId:   in.Id,
		AggregateType: "Payment",
		Version:       1,
		Payload:       string(payload),
	}, nil
}

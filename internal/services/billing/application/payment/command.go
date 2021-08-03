package payment_application

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"google.golang.org/protobuf/encoding/protojson"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

func CommandPaymentCreate(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	aggregateId := uuid.New().String()
	in.Status = billing.StatusPayment_STATUS_PAYMENT_NEW
	in.Id = aggregateId

	// start tracing
	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("command: UpdateBalance"))
	span.SetTag("aggregate id", aggregateId)
	span.SetTag("command type", billing.Command_COMMAND_PAYMENT_CREATE.String())
	defer span.Finish()

	payload, err := protojson.Marshal(in)
	if err != nil {
		span.SetTag("error", true)
		span.SetTag("message", err.Error())
		return nil, err
	}

	span.LogFields(opentracinglog.String("log", string(payload)))

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
	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("command: UpdateBalance"))
	span.SetTag("aggregate id", in.Id)
	span.SetTag("command type", billing.Command_COMMAND_BALANCE_UPDATE.String())
	defer span.Finish()

	payload, err := protojson.Marshal(in)
	if err != nil {
		span.SetTag("error", true)
		span.SetTag("message", err.Error())
		return nil, err
	}

	span.LogFields(opentracinglog.String("log", string(payload)))

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
	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("command: PaymentClose"))
	span.SetTag("aggregate id", in.Id)
	span.SetTag("command type", billing.Command_COMMAND_PAYMENT_CLOSE.String())
	defer span.Finish()

	payload, err := protojson.Marshal(in)
	if err != nil {
		span.SetTag("error", true)
		span.SetTag("message", err.Error())
		return nil, err
	}

	span.LogFields(opentracinglog.String("log", string(payload)))

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_CLOSE.String(),
		AggregateId:   in.Id,
		AggregateType: "Payment",
		Version:       1,
		Payload:       string(payload),
	}, nil
}

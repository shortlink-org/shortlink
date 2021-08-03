package payment_application

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

func CommandPaymentCreate(in *Payment) (*eventsourcing.BaseCommand, error) {
	in.Status = billing.StatusPayment_STATUS_PAYMENT_NEW

	payload, err := protojson.Marshal(in.Payment)
	if err != nil {
		return nil, err
	}

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_CREATE.String(),
		AggregateId:   uuid.New().String(),
		AggregateType: "Payment",
		Version:       0,
		Payload:       string(payload),
	}, nil
}

func CommandPaymentUpdateBalance(in *Payment) (*eventsourcing.BaseCommand, error) {
	payload, err := protojson.Marshal(in.Payment)
	if err != nil {
		return nil, err
	}

	return &eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_BALANCE_UPDATE.String(),
		AggregateId:   in.Payment.Id,
		AggregateType: "Payment",
		Version:       in.Version,
		Payload:       string(payload),
	}, nil
}

package payment_application

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	"github.com/batazor/shortlink/internal/pkg/logger"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
	billing_store "github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
)

type PaymentService struct {
	logger logger.Logger

	// Repositories
	paymentRepository billing_store.PaymentRepository
}

func New(logger logger.Logger, paymentRepository billing_store.PaymentRepository) (*PaymentService, error) {
	return &PaymentService{
		logger: logger,

		// Repositories
		paymentRepository: paymentRepository,
	}, nil
}

func (p *PaymentService) Get(ctx context.Context, id string) (*billing.Payment, error) {
	panic("implement me")
	//return p.paymentRepository.Get(ctx, id)
}

func (p *PaymentService) List(ctx context.Context, filter interface{}) ([]*billing.Payment, error) {
	panic("implement me")
	//return p.paymentRepository.List(ctx, filter)
}

// Add - Create a payment
func (p *PaymentService) Add(ctx context.Context, in *billing.Payment) (*billing.Payment, error) {
	payment := &Payment{
		Payment:       in,
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	payload, err := protojson.Marshal(in)
	if err != nil {
		return nil, err
	}

	err = payment.HandleCommand(&eventsourcing.BaseCommand{
		Type:          billing.Command_COMMAND_PAYMENT_CREATE.String(),
		AggregateId:   uuid.New().String(),
		AggregateType: "Payment",
		Version:       0,
		Payload:       string(payload),
	})
	if err != nil {
		return nil, err
	}

	//payment.
	//resp, err := p.paymentRepository.Add(ctx, payment)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}

func (p *PaymentService) Update(ctx context.Context, in *billing.Payment) (*billing.Payment, error) {
	panic("implement me")
	//return p.paymentRepository.Update(ctx, in)
}

func (p *PaymentService) Delete(ctx context.Context, id string) error {
	panic("implement me")
	//return p.paymentRepository.Delete(ctx, id)
}

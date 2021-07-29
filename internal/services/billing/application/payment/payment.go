package payment_application

import (
	"context"

	event_store "github.com/batazor/shortlink/internal/pkg/eventsourcing/store"
	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
	"github.com/batazor/shortlink/internal/pkg/logger"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

type PaymentService struct {
	logger logger.Logger

	// Repositories
	paymentRepository event_store.EventStore
}

func New(logger logger.Logger, paymentRepository event_store.EventStore) (*PaymentService, error) {
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
	aggregate := &Payment{
		Payment:       in,
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentCreate(aggregate)
	if err != nil {
		return nil, err
	}

	err = aggregate.HandleCommand(command)
	if err != nil {
		return nil, err
	}

	err = p.paymentRepository.Save(ctx, aggregate.Uncommitted())
	if err != nil {
		return nil, err
	}

	// TODO: PublishEvent(aggregate)

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

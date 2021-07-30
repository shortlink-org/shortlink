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

	// EventSourcing
	eventsourcing.CommandHandle

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

func (p *PaymentService) Get(ctx context.Context, aggregateId string) (*billing.Payment, error) {
	aggregate := &Payment{
		Payment:       &billing.Payment{},
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	events, err := p.paymentRepository.Load(ctx, aggregateId)
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		errApplyChange := aggregate.ApplyChangeHelper(aggregate, event, false)
		if errApplyChange != nil {
			return nil, errApplyChange
		}
	}

	return aggregate.Payment, nil
}

func (p *PaymentService) List(ctx context.Context, filter interface{}) ([]*billing.Payment, error) {
	panic("implement me")
	//return p.paymentRepository.List(ctx, filter)
}

func (p *PaymentService) Handle(ctx context.Context, in *billing.Payment) (*eventsourcing.BaseCommand, error) {
	aggregate := &Payment{
		Payment:       in,
		BaseAggregate: &eventsourcing.BaseAggregate{},
	}

	command, err := CommandPaymentCreate(aggregate)
	if err != nil {
		return nil, err
	}

	// Check update or create
	if command.Version > 1 {
		events, errLoad := p.paymentRepository.Load(ctx, command.AggregateId)
		if errLoad != nil {
			return nil, errLoad
		}

		for _, event := range events {
			errApplyChange := aggregate.ApplyChangeHelper(aggregate, event, false)
			if errApplyChange != nil {
				return nil, errApplyChange
			}
		}
	}

	err = aggregate.HandleCommand(command)
	if err != nil {
		return nil, err
	}

	err = p.paymentRepository.Save(ctx, aggregate.Uncommitted())
	if err != nil {
		return nil, err
	}

	return command, nil
}

// Add - Create a payment
func (p *PaymentService) Add(ctx context.Context, in *billing.Payment) (*billing.Payment, error) {
	command, err := p.Handle(ctx, in)
	if err != nil {
		return nil, err
	}

	// safe identity
	in.Id = command.AggregateId

	// TODO: PublishEvent(aggregate)
	return in, nil
}

func (p *PaymentService) Update(ctx context.Context, in *billing.Payment) (*billing.Payment, error) {
	_, err := p.Handle(ctx, in)
	if err != nil {
		return nil, err
	}

	// TODO: PublishEvent(aggregate)
	return in, nil
}

func (p *PaymentService) Delete(ctx context.Context, id string) error {
	panic("implement me")
	//return p.paymentRepository.Delete(ctx, id)
}

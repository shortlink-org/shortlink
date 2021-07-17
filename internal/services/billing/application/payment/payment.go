package payment_application

import (
	"context"

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

func (p PaymentService) Get(ctx context.Context, id string) (*billing.Payment, error) {
	return p.paymentRepository.Get(ctx, id)
}

func (p PaymentService) List(ctx context.Context, filter interface{}) ([]*billing.Payment, error) {
	return p.paymentRepository.List(ctx, filter)
}

func (p PaymentService) Add(ctx context.Context, in *billing.Payment) (*billing.Payment, error) {
	return p.paymentRepository.Add(ctx, in)
}

func (p PaymentService) Update(ctx context.Context, in *billing.Payment) (*billing.Payment, error) {
	return p.paymentRepository.Update(ctx, in)
}

func (p PaymentService) Delete(ctx context.Context, id string) error {
	return p.paymentRepository.Delete(ctx, id)
}

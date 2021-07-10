package balance_application

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/logger"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
	billing_store "github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
)

type BalanceService struct {
	logger logger.Logger

	// Repositories
	balanceRepository billing_store.BalanceRepository
}

func New(logger logger.Logger, balanceRepository billing_store.BalanceRepository) (*BalanceService, error) {
	return &BalanceService{
		logger: logger,

		// Repositories
		balanceRepository: balanceRepository,
	}, nil
}

func (b *BalanceService) Get(ctx context.Context, id *billing.Balance) (*billing.Balance, error) {
	return b.balanceRepository.Get(ctx, id)
}

func (b *BalanceService) List(ctx context.Context, filter interface{}) ([]*billing.Balance, error) {
	return b.balanceRepository.List(ctx, filter)
}

func (b *BalanceService) Update(ctx context.Context, in *billing.Balance) (*billing.Balance, error) {
	return b.balanceRepository.Update(ctx, in)
}

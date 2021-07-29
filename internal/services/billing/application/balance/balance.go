package balance_application

import (
	"context"

	event_store "github.com/batazor/shortlink/internal/pkg/eventsourcing/store"
	"github.com/batazor/shortlink/internal/pkg/logger"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/balance/v1"
)

type BalanceService struct {
	logger logger.Logger

	// Repositories
	balanceRepository event_store.EventStore
}

func New(logger logger.Logger, balanceRepository event_store.EventStore) (*BalanceService, error) {
	return &BalanceService{
		logger: logger,

		// Repositories
		balanceRepository: balanceRepository,
	}, nil
}

func (b *BalanceService) Get(ctx context.Context, id *billing.Balance) (*billing.Balance, error) {
	//return b.balanceRepository.Get(ctx, id)
	panic("implement me")
}

func (b *BalanceService) List(ctx context.Context, filter interface{}) ([]*billing.Balance, error) {
	//return b.balanceRepository.List(ctx, filter)
	panic("implement me")
}

func (b *BalanceService) Update(ctx context.Context, in *billing.Balance) (*billing.Balance, error) {
	//return b.balanceRepository.Update(ctx, in)
	panic("implement me")
}

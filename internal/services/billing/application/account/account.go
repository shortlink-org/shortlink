package account_application

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/logger"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
	billing_store "github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
)

type AccountService struct {
	logger logger.Logger

	// Repositories
	accountRepository billing_store.AccountRepository
}

func New(logger logger.Logger, accountRepository billing_store.AccountRepository) (*AccountService, error) {
	return &AccountService{
		logger: logger,

		// Repositories
		accountRepository: accountRepository,
	}, nil
}

func (acc *AccountService) Add(ctx context.Context, in *billing.Account) (*billing.Account, error) {
	return acc.accountRepository.Add(ctx, in)
}

package account_application

import (
	"context"

	"github.com/google/uuid"

	billing "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/domain/billing/account/v1"
	billing_store "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/infrastructure/store"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

type AccountService struct {
	log logger.Logger

	// Repositories
	accountRepository billing_store.AccountRepository
}

func New(log logger.Logger, accountRepository billing_store.AccountRepository) (*AccountService, error) {
	return &AccountService{
		log: log,

		// Repositories
		accountRepository: accountRepository,
	}, nil
}

func (acc *AccountService) Get(ctx context.Context, id string) (*billing.Account, error) {
	return acc.accountRepository.Get(ctx, id)
}

func (acc *AccountService) List(ctx context.Context, filter any) ([]*billing.Account, error) {
	return acc.accountRepository.List(ctx, filter)
}

func (acc *AccountService) Add(ctx context.Context, in *billing.Account) (*billing.Account, error) {
	// generate uniq identity
	in.Id = uuid.New().String()

	return acc.accountRepository.Add(ctx, in)
}

func (acc *AccountService) Update(ctx context.Context, in *billing.Account) (*billing.Account, error) {
	return acc.accountRepository.Update(ctx, in)
}

func (acc *AccountService) Delete(ctx context.Context, id string) error {
	return acc.accountRepository.Delete(ctx, id)
}

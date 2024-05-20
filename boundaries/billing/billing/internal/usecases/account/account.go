package account_application

import (
	"context"

	domain "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1"
	"github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1/rules"
	account_repository "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/repository/account"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/notify"
	"github.com/shortlink-org/shortlink/pkg/pattern/specification"
)

type AccountService struct {
	log logger.Logger

	// Subscriber
	// notify.Subscriber[link.Link]

	// Repositories
	accountRepository account_repository.Repository
}

func New(ctx context.Context, log logger.Logger, conn db.DB) (*AccountService, error) {
	// Init Repository
	accountRepository, err := account_repository.New(ctx, conn)
	if err != nil {
		return nil, err
	}

	service := &AccountService{
		log: log,

		// Repositories
		accountRepository: accountRepository,
	}

	// Subscribe to Event ==============================================================================================
	notify.Subscribe(uint32(domain.Event_EVENT_ACCOUNT_NEW), service)
	notify.Subscribe(uint32(domain.Event_EVENT_ACCOUNT_DELETE), service)

	log.Info("init usecase", field.Fields{
		"name": "account",
	})

	return service, nil
}

// Notify - implementation of notify.Subscriber interface
func (acc *AccountService) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	return notify.Response[any]{}
}

func (acc *AccountService) Get(ctx context.Context, id string) (*domain.Account, error) {
	return acc.accountRepository.Get(ctx, id)
}

func (acc *AccountService) List(ctx context.Context, filter any) ([]*domain.Account, error) {
	return acc.accountRepository.List(ctx, filter)
}

func (acc *AccountService) Add(ctx context.Context, in *domain.Account) (*domain.Account, error) {
	// create specification
	spec := specification.NewAndSpecification[domain.Account](
		rules.NewUserId(),
		rules.NewTariffId(),
	)

	err := spec.IsSatisfiedBy(in)
	if err != nil {
		return nil, err
	}

	return acc.accountRepository.Add(ctx, in)
}

func (acc *AccountService) Update(ctx context.Context, in *domain.Account) (*domain.Account, error) {
	return acc.accountRepository.Update(ctx, in)
}

func (acc *AccountService) Delete(ctx context.Context, id string) error {
	return acc.accountRepository.Delete(ctx, id)
}

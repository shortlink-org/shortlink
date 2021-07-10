package billing_store

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/notify"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
)

// Store abstract type
type BillingStore struct {
	typeStore string

	// Repositories
	Account AccountRepository
	Order   OrderRepository
	Payment PaymentRepository
	Balance BalanceRepository
	Tariff  TariffRepository

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}

type Repository interface {
	Init(ctx context.Context, db *db.Store) error
}

type AccountRepository interface {
	Repository

	Get(ctx context.Context, id string) (*billing.Account, error)
	List(ctx context.Context, filter interface{}) ([]*billing.Account, error)
	Add(ctx context.Context, in *billing.Account) (*billing.Account, error)
	Update(ctx context.Context, in *billing.Account) (*billing.Account, error)
	Delete(ctx context.Context, id string) error
}

type OrderRepository interface {
	Repository

	Get(ctx context.Context, id string) (*billing.Order, error)
	List(ctx context.Context, filter interface{}) ([]*billing.Order, error)
	Add(ctx context.Context, in *billing.Order) (*billing.Order, error)
	Update(ctx context.Context, in *billing.Order) (*billing.Order, error)
	Delete(ctx context.Context, id string) error
}

type PaymentRepository interface {
	Repository

	Get(ctx context.Context, id string) (*billing.Payment, error)
	List(ctx context.Context, filter interface{}) ([]*billing.Payment, error)
	Add(ctx context.Context, in *billing.Payment) (*billing.Payment, error)
	Update(ctx context.Context, in *billing.Payment) (*billing.Payment, error)
	Delete(ctx context.Context, id string) error
}

type BalanceRepository interface {
	Repository

	Get(ctx context.Context, id *billing.Balance) (*billing.Balance, error)
	List(ctx context.Context, filter interface{}) ([]*billing.Balance, error)
	Update(ctx context.Context, in *billing.Balance) (*billing.Balance, error)
}

type TariffRepository interface {
	Repository

	Get(ctx context.Context, id string) (*billing.Tariff, error)
	List(ctx context.Context, filter interface{}) ([]*billing.Tariff, error)
	Add(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Update(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Delete(ctx context.Context, id string) error
}

package billing_store

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/billing/domain/billing/account/v1"
	v12 "github.com/batazor/shortlink/internal/services/billing/domain/billing/balance/v1"
	v13 "github.com/batazor/shortlink/internal/services/billing/domain/billing/order/v1"
	v14 "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/tariff/v1"
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

	Get(ctx context.Context, id string) (*v1.Account, error)
	List(ctx context.Context, filter interface{}) ([]*v1.Account, error)
	Add(ctx context.Context, in *v1.Account) (*v1.Account, error)
	Update(ctx context.Context, in *v1.Account) (*v1.Account, error)
	Delete(ctx context.Context, id string) error
}

type OrderRepository interface {
	Repository

	Get(ctx context.Context, id string) (*v13.Order, error)
	List(ctx context.Context, filter interface{}) ([]*v13.Order, error)
	Add(ctx context.Context, in *v13.Order) (*v13.Order, error)
	Update(ctx context.Context, in *v13.Order) (*v13.Order, error)
	Delete(ctx context.Context, id string) error
}

type PaymentRepository interface {
	Repository

	Get(ctx context.Context, id string) (*v14.Payment, error)
	List(ctx context.Context, filter interface{}) ([]*v14.Payment, error)
	Add(ctx context.Context, in *v14.Payment) (*v14.Payment, error)
	Update(ctx context.Context, in *v14.Payment) (*v14.Payment, error)
	Delete(ctx context.Context, id string) error
}

type BalanceRepository interface {
	Repository

	Get(ctx context.Context, id *v12.Balance) (*v12.Balance, error)
	List(ctx context.Context, filter interface{}) ([]*v12.Balance, error)
	Update(ctx context.Context, in *v12.Balance) (*v12.Balance, error)
}

type TariffRepository interface {
	Repository

	Get(ctx context.Context, id string) (*billing.Tariff, error)
	List(ctx context.Context, filter interface{}) (*billing.Tariffs, error)
	Add(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Update(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Delete(ctx context.Context, id string) error
}

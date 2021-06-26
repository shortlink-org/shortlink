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
	Store     Repository
	Account   AccountRepository
	Tariff    TariffRepository

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}

type Repository interface {
	Init(ctx context.Context, db *db.Store) error
}

type AccountRepository interface {
	Init(ctx context.Context, db *db.Store) error

	Get(ctx context.Context, id string) (*billing.Account, error)
	List(ctx context.Context, filter interface{}) (*billing.Account, error)
	Add(ctx context.Context, in *billing.Account) (*billing.Account, error)
	Update(ctx context.Context, in *billing.Account) (*billing.Account, error)
	Delete(ctx context.Context, id string) error
}

type TariffRepository interface {
	Init(ctx context.Context, db *db.Store) error

	Get(ctx context.Context, id string) (*billing.Tariff, error)
	List(ctx context.Context, filter interface{}) (*billing.Tariff, error)
	Add(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Update(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Delete(ctx context.Context, id string) error
}

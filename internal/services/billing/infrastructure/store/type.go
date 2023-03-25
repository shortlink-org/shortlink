package billing_store

import (
	"context"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	event_store "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/store"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	v1 "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/account/v1"
	billing "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/tariff/v1"
	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// Store abstract type
type BillingStore struct {
	typeStore string

	// Repositories
	Account    AccountRepository
	Tariff     TariffRepository
	EventStore *event_store.Repository

	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]
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

type TariffRepository interface {
	Repository

	Get(ctx context.Context, id string) (*billing.Tariff, error)
	List(ctx context.Context, filter interface{}) (*billing.Tariffs, error)
	Add(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Update(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Delete(ctx context.Context, id string) error
}

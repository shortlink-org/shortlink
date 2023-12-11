package billing_store

import (
	"context"

	link "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	v1 "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/domain/billing/account/v1"
	billing "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/domain/billing/tariff/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	event_store "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/store"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

// Store abstract type
type BillingStore struct {
	Account AccountRepository
	Tariff  TariffRepository
	notify.Subscriber[link.Link]
	EventStore *event_store.Repository
	typeStore  string
}

type Repository interface {
	Init(ctx context.Context, store db.DB) error
}

type AccountRepository interface {
	Repository

	Get(ctx context.Context, id string) (*v1.Account, error)
	List(ctx context.Context, filter any) ([]*v1.Account, error)
	Add(ctx context.Context, in *v1.Account) (*v1.Account, error)
	Update(ctx context.Context, in *v1.Account) (*v1.Account, error)
	Delete(ctx context.Context, id string) error
}

type TariffRepository interface {
	Repository

	Get(ctx context.Context, id string) (*billing.Tariff, error)
	List(ctx context.Context, filter any) (*billing.Tariffs, error)
	Add(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Update(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error)
	Delete(ctx context.Context, id string) error
}

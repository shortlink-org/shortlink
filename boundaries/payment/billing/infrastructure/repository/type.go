package billing_store

import (
	link "github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1"
	account_repository "github.com/shortlink-org/shortlink/boundaries/payment/billing/infrastructure/repository/account"
	tariff_repository "github.com/shortlink-org/shortlink/boundaries/payment/billing/infrastructure/repository/tariff"
	"github.com/shortlink-org/shortlink/pkg/eventsourcing"
	"github.com/shortlink-org/shortlink/pkg/notify"
)

// Store - billing store
type Store struct {
	notify.Subscriber[link.Link]
	EventSourcing *eventsourcing.EventSourcing
	typeStore     string

	// Repositories
	Account account_repository.Repository
	Tariff  tariff_repository.Repository
}

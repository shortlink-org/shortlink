/*
Billing Service. Infrastructure layer
*/
package billing_store

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	event_store "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/store"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	account "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/account/v1"
	tariff "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/tariff/v1"
	"github.com/shortlink-org/shortlink/internal/services/billing/infrastructure/store/postgres"
)

// Use return implementation of db
func (s *BillingStore) Use(ctx context.Context, log logger.Logger, db *db.Store) (*BillingStore, error) {
	// Set configuration
	s.setConfig()

	// Subscribe to Event ==============================================================================================
	// account events
	notify.Subscribe(uint32(account.Event_EVENT_ACCOUNT_NEW), s)
	notify.Subscribe(uint32(account.Event_EVENT_ACCOUNT_DELETE), s)

	// tariff events
	notify.Subscribe(uint32(tariff.Event_EVENT_TARIFF_NEW), s)
	notify.Subscribe(uint32(tariff.Event_EVENT_TARIFF_UPDATE), s)
	notify.Subscribe(uint32(tariff.Event_EVENT_TARIFF_CLOSE), s)

	switch s.typeStore {
	case "postgres":
		s.Account = &postgres.Account{}
		s.Tariff = &postgres.Tariff{}
		s.EventStore = &event_store.Repository{}
	default:
		s.Account = &postgres.Account{}
		s.Tariff = &postgres.Tariff{}
		s.EventStore = &event_store.Repository{}
	}

	err := s.Account.Init(ctx, db)
	if err != nil {
		return nil, err
	}

	err = s.Tariff.Init(ctx, db)
	if err != nil {
		return nil, err
	}

	_, err = s.EventStore.Use(ctx, log, db)
	if err != nil {
		return nil, err
	}

	log.Info("init billingStore", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

// Notify - implementation of notify.Subscriber interface
func (s *BillingStore) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	return notify.Response[any]{}
}

func (s *BillingStore) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}

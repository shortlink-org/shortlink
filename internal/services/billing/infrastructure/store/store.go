/*
Billing Service. Infrastructure layer
*/
package billing_store

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/notify"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/store/postgres"
)

// Use return implementation of db
func (s *BillingStore) Use(ctx context.Context, log logger.Logger, db *db.Store) (*BillingStore, error) {
	// Set configuration
	s.setConfig()

	// Subscribe to Event ==============================================================================================
	// account events
	notify.Subscribe(uint32(billing.Event_EVENT_ACCOUNT_NEW), s)
	notify.Subscribe(uint32(billing.Event_EVENT_ACCOUNT_DELETE), s)

	// tariff events
	notify.Subscribe(uint32(billing.Event_EVENT_TARIFF_NEW), s)
	notify.Subscribe(uint32(billing.Event_EVENT_TARIFF_UPDATE), s)
	notify.Subscribe(uint32(billing.Event_EVENT_TARIFF_CLOSE), s)

	switch s.typeStore {
	case "postgres":
		s.Account = &postgres.Account{}
		s.Balance = &postgres.Balance{}
		s.Order = &postgres.Order{}
		s.Payment = &postgres.Payment{}
		s.Tariff = &postgres.Tariff{}
	default:
		s.Account = &postgres.Account{}
		s.Balance = &postgres.Balance{}
		s.Order = &postgres.Order{}
		s.Payment = &postgres.Payment{}
		s.Tariff = &postgres.Tariff{}
	}

	_ = s.Account.Init(ctx, db)
	_ = s.Balance.Init(ctx, db)
	_ = s.Order.Init(ctx, db)
	_ = s.Payment.Init(ctx, db)
	_ = s.Tariff.Init(ctx, db)

	log.Info("init billingStore", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

func (s *BillingStore) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}

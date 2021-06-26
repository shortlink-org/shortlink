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
		s.Store = &postgres.Store{}
	default:
		s.Store = &postgres.Store{}
	}
	_ = s.Store.Init(ctx, db)

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

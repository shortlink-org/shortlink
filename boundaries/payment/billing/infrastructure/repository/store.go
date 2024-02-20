/*
Billing Service. Infrastructure layer
*/
package billing_store

import (
	"context"

	"github.com/spf13/viper"

	account "github.com/shortlink-org/shortlink/boundaries/payment/billing/domain/billing/account/v1"
	tariff "github.com/shortlink-org/shortlink/boundaries/payment/billing/domain/billing/tariff/v1"
	account_repository "github.com/shortlink-org/shortlink/boundaries/payment/billing/infrastructure/repository/account"
	tariff_repository "github.com/shortlink-org/shortlink/boundaries/payment/billing/infrastructure/repository/tariff"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/notify"
)

// Use return implementation of db
func (s *Store) Use(ctx context.Context, log logger.Logger, store db.DB) (*Store, error) {
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

	// Init Repository =================================================================================================
	var err error

	s.Account, err = account_repository.New(ctx, store)
	if err != nil {
		return nil, err
	}

	s.Tariff, err = tariff_repository.New(ctx, store)
	if err != nil {
		return nil, err
	}

	log.Info("init billingStore", field.Fields{
		"store": s.typeStore,
	})

	return s, nil
}

// Notify - implementation of notify.Subscriber interface
func (s *Store) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	return notify.Response[any]{}
}

func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}

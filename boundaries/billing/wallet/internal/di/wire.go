//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Billing Service DI-package
*/
package wallet_di

import (
	"github.com/google/wire"

	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

type WalletService struct {
	Log logger.Logger
}

// WalletService =======================================================================================================
var WalletSet = wire.NewSet(
	di.DefaultSet,

	NewWalletService,
)

func NewWalletService(log logger.Logger) (*WalletService, error) {
	log.Info("Start wallet service")

	return &WalletService{
		Log: log,
	}, nil
}

func InitializeWalletService() (*WalletService, func(), error) {
	panic(wire.Build(WalletSet))
}

//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Billing Service DI-package
*/
package wallet_di

import (
	"github.com/google/wire"

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

type WalletService struct {
	Logger logger.Logger
}

// WalletService =======================================================================================================
var WalletSet = wire.NewSet(
	di.DefaultSet,

	NewWalletService,
)

func NewWalletService(logger logger.Logger) (*WalletService, error) {
	logger.Info("Start wallet service")

	return &WalletService{
		Logger: logger,
	}, nil
}

func InitializeWalletService() (*WalletService, func(), error) {
	panic(wire.Build(WalletSet))
}

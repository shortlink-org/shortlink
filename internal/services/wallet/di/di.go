//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wallet_di

import (
	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/pkg/logger"
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

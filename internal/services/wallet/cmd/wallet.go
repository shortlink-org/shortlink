/*
Wallet application

Wallet-service
*/
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/services/wallet/di"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-wallet")

	// Init a new service
	service, cleanup, err := wallet_di.InitializeWalletService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Logger.Error(r.(string))
		}
	}()

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	cleanup()
}

package wallet

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type WalletService struct {
	client *ethclient.Client
}

func NewWalletService() (*WalletService, error) {
	var err error
	wallet := &WalletService{}

	wallet.client, err = NewWalletClient()
	if err != nil {
		return nil, err
	}

	fmt.Println("we have a connection")
	return wallet, nil
}

func NewWalletClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		return nil, err
	}

	return client, nil
}

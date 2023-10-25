package wallet

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Service struct {
	client *ethclient.Client
}

func NewService() (*Service, error) {
	var err error
	wallet := &Service{}

	wallet.client, err = NewClient()
	if err != nil {
		return nil, err
	}

	//nolint:revive // TODO: add logger later
	fmt.Println("we have a connection")

	return wallet, nil
}

func NewClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		return nil, err
	}

	return client, nil
}

package rules

import (
	"errors"

	"github.com/google/uuid"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1"
)

var ErrTariffIdRequired = errors.New("tariffId is required")

type TariffId struct{}

func NewTariffId() *TariffId {
	return &TariffId{}
}

func (t *TariffId) IsSatisfiedBy(account *v1.Account) error {
	if account.GetTariffId() != uuid.Nil {
		return nil
	}

	return ErrTariffIdRequired
}

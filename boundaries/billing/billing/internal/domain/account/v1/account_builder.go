package v1

import (
	"errors"

	"github.com/google/uuid"
)

// AccountBuilder is used to build a new Account
type AccountBuilder struct {
	account *Account
	errors  error
}

// NewAccountBuilder returns a new instance of AccountBuilder
func NewAccountBuilder() *AccountBuilder {
	return &AccountBuilder{account: &Account{}}
}

// SetId sets the id of the account
func (b *AccountBuilder) SetId(id uuid.UUID) *AccountBuilder {
	if id == uuid.Nil {
		b.errors = errors.Join(b.errors, ErrInvalidAccountId)
		return b
	}

	b.account.id = id

	return b
}

// SetUserId sets the userId of the account
func (b *AccountBuilder) SetUserId(userId uuid.UUID) *AccountBuilder {
	if userId == uuid.Nil {
		b.errors = errors.Join(b.errors, ErrInvalidAccountUserId)
		return b
	}

	b.account.userId = userId

	return b
}

// SetTariffId sets the tariffId of the account
func (b *AccountBuilder) SetTariffId(tariffId uuid.UUID) *AccountBuilder {
	if tariffId == uuid.Nil {
		b.errors = errors.Join(b.errors, ErrInvalidAccountTariffId)
		return b
	}

	b.account.tariffId = tariffId

	return b
}

// Build finalizes the building process and returns the built Account
func (b *AccountBuilder) Build() (*Account, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	// Generate a new id if it is not set
	if b.account.id == uuid.Nil {
		var err error

		b.account.id, err = uuid.NewV7()
		if err != nil {
			return nil, err
		}
	}

	return b.account, nil
}

package rules

import (
	"errors"

	"github.com/google/uuid"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1"
)

var ErrUserIdRequired = errors.New("userId is required")

type UserId struct{}

func NewUserId() *UserId {
	return &UserId{}
}

func (u *UserId) IsSatisfiedBy(account *v1.Account) error {
	if account.GetUserId() != uuid.Nil {
		return nil
	}

	return ErrUserIdRequired
}

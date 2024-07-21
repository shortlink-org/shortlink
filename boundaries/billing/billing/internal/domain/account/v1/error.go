package v1

import (
	"errors"
)

var (
	ErrInvalidAccountId       = errors.New("invalid id: id is nil")
	ErrInvalidAccountUserId   = errors.New("invalid userId: userId is nil")
	ErrInvalidAccountTariffId = errors.New("invalid tariffId: tariffId is nil")
)

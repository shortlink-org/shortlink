package v1

import (
	"errors"
)

var (
	ErrInvalidOrderId       = errors.New("invalid id: id is empty")
	ErrInvalidOrderUserId   = errors.New("invalid userId: userId is empty")
	ErrInvalidOrderTariffId = errors.New("invalid tariffId: tariffId is empty")
	ErrInvalidOrderStatus   = errors.New("invalid status: status is not recognized")
)

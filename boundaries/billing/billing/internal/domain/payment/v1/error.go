package v1

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPaymentId     = errors.New("invalid id: id is empty")
	ErrInvalidPaymentName   = errors.New("invalid name: name is empty")
	ErrInvalidPaymentStatus = errors.New("invalid status: status is not recognized")
	ErrInvalidPaymentUserId = errors.New("invalid userId: userId is empty")
	ErrInvalidPaymentAmount = errors.New("invalid amount: amount must be greater than zero")
)

// IncorrectStatusOfPaymentError is an error type for incorrect payment statuses
type IncorrectStatusOfPaymentError struct {
	Status string
}

// Error implements the error interface for IncorrectStatusOfPaymentError
func (e *IncorrectStatusOfPaymentError) Error() string {
	return fmt.Sprintf("incorrect status of payment: %s", e.Status)
}

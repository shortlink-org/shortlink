package v1

import (
	"fmt"
)

// IncorrectStatusOfPaymentError is an error type for incorrect payment statuses
type IncorrectStatusOfPaymentError struct {
	Status string
}

// Error implements the error interface for IncorrectStatusOfPaymentError
func (e *IncorrectStatusOfPaymentError) Error() string {
	return fmt.Sprintf("incorrect status of payment: %s", e.Status)
}

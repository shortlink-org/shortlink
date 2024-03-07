package payment_application

import (
	"fmt"
)

var ErrCreatePayment = fmt.Errorf("error create a new payment")

var ErrApprovePayment = fmt.Errorf("Payment was successfully created, but its status could not be received")

type IncorrectStatusOfPaymentError struct {
	Status string
}

func (e *IncorrectStatusOfPaymentError) Error() string {
	return fmt.Sprintf("incorrect status of payment: %s", e.Status)
}

type NotFoundEventError struct {
	Type string
}

func (e *NotFoundEventError) Error() string {
	return fmt.Sprintf("not found event with type: %s", e.Type)
}

type NotFoundCommandError struct {
	Type string
}

func (e *NotFoundCommandError) Error() string {
	return fmt.Sprintf("not found command with type: %s", e.Type)
}

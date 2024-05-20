package v1

import (
	"github.com/google/uuid"
)

// PAYMENT =============================================================================================================

// EventPaymentCreated is published when a payment is created
type EventPaymentCreated struct {
	// id of the payment
	Id uuid.UUID `json:"id,omitempty"`
	// name of the payment
	Name string `json:"name,omitempty"`
	// status of the payment
	Status StatusPayment `json:"status,omitempty"`
	// owner of the payment
	UserId uuid.UUID `json:"user_id,omitempty"`
}

// EventPaymentApproved is published when a payment is approved
type EventPaymentApproved struct {
	// id of the payment
	Id uuid.UUID `json:"id,omitempty"`
	// status of the payment
	Status StatusPayment `json:"status,omitempty"`
}

// EventPaymentRejected is published when a payment is rejected
type EventPaymentRejected struct {
	// id of the payment
	Id uuid.UUID `json:"id,omitempty"`
	// status of the payment
	Status StatusPayment `json:"status,omitempty"`
}

// EventPaymentClosed is published when a payment is closed
type EventPaymentClosed struct {
	// id of the payment
	Id uuid.UUID `json:"id,omitempty"`
	// status of the payment
	Status StatusPayment `json:"status,omitempty"`
}

// BALANCE =============================================================================================================

// EventBalanceUpdated is published when a balance is updated
type EventBalanceUpdated struct {
	// id of the balance
	Id uuid.UUID `json:"id,omitempty"`
	// amount of the balance
	Amount int64 `json:"amount,omitempty"`
}

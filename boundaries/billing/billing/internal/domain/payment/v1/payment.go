package v1

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Payment - information about payment
type Payment struct {
	// id payment
	id uuid.UUID
	// name payment
	name string
	// status payment
	status StatusPayment
	// User ID
	userId uuid.UUID
	// Amount payment
	amount decimal.Decimal
}

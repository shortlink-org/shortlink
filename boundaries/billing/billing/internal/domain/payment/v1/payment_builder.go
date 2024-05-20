package v1

import (
	"errors"

	"github.com/google/uuid"
)

// PaymentBuilder is used to build a new Payment
type PaymentBuilder struct {
	payment *Payment
	errors  error
}

// NewPaymentBuilder returns a new instance of PaymentBuilder
func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{payment: &Payment{}}
}

// SetId sets the id of the payment
func (b *PaymentBuilder) SetId(id uuid.UUID) *PaymentBuilder {
	if id == uuid.Nil {
		b.errors = errors.Join(b.errors, errors.New("invalid id: id is empty"))
		return b
	}

	b.payment.id = id
	return b
}

// SetName sets the name of the payment
func (b *PaymentBuilder) SetName(name string) *PaymentBuilder {
	if name == "" {
		b.errors = errors.Join(b.errors, errors.New("invalid name: name is empty"))
		return b
	}

	b.payment.name = name
	return b
}

// SetStatus sets the status of the payment
func (b *PaymentBuilder) SetStatus(status StatusPayment) *PaymentBuilder {
	if _, ok := StatusPayment_name[int32(status)]; !ok {
		b.errors = errors.Join(b.errors, errors.New("invalid status: status is not recognized"))
		return b
	}

	b.payment.status = status
	return b
}

// SetUserId sets the userId of the payment
func (b *PaymentBuilder) SetUserId(userId uuid.UUID) *PaymentBuilder {
	if userId == uuid.Nil {
		b.errors = errors.Join(b.errors, errors.New("invalid userId: userId is empty"))
		return b
	}

	b.payment.userId = userId
	return b
}

// SetAmount sets the amount of the payment
func (b *PaymentBuilder) SetAmount(amount int64) *PaymentBuilder {
	if amount <= 0 {
		b.errors = errors.Join(b.errors, errors.New("invalid amount: amount must be greater than zero"))
		return b
	}

	b.payment.amount = amount
	return b
}

// Build finalizes the building process and returns the built Payment
func (b *PaymentBuilder) Build() (*Payment, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	return b.payment, nil
}

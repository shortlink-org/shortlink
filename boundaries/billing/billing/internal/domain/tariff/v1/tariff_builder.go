package v1

import (
	"errors"
)

// TariffBuilder is used to build a new Tariff
type TariffBuilder struct {
	tariff *Tariff
	errors error
}

// NewTariffBuilder returns a new instance of TariffBuilder
func NewTariffBuilder() *TariffBuilder {
	return &TariffBuilder{tariff: &Tariff{}}
}

// SetId sets the id of the tariff
func (b *TariffBuilder) SetId(id string) *TariffBuilder {
	if id == "" {
		b.errors = errors.Join(b.errors, ErrInvalidId)
		return b
	}

	b.tariff.id = id

	return b
}

// SetName sets the name of the tariff
func (b *TariffBuilder) SetName(name string) *TariffBuilder {
	if name == "" {
		b.errors = errors.Join(b.errors, ErrInvalidName)
		return b
	}

	b.tariff.name = name

	return b
}

// SetPayload sets the payload of the tariff
func (b *TariffBuilder) SetPayload(payload string) *TariffBuilder {
	if payload == "" {
		b.errors = errors.Join(b.errors, ErrInvalidPayload)

		return b
	}

	b.tariff.payload = payload

	return b
}

// Build finalizes the building process and returns the built Tariff
func (b *TariffBuilder) Build() (*Tariff, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	return b.tariff, nil
}

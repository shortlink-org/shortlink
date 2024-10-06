package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CartItem struct {
	ProductID uuid.UUID
	Quantity  int32
	Price     decimal.Decimal
	Brand     string
}

type Cart struct {
	Items      []CartItem
	CustomerID uuid.UUID
}

func (c *Cart) AddItem(item CartItem) {
	c.Items = append(c.Items, item)
}

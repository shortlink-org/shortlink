package v1

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Items represent a list of order items.
type Items []Item

// Item represents an item in the order.
type Item struct {
	productId uuid.UUID
	quantity  int
	price     decimal.Decimal
}

// NewItem creates a new item.
func NewItem(productId uuid.UUID, quantity int, price decimal.Decimal) Item {
	return Item{
		productId: productId,
		quantity:  quantity,
		price:     price,
	}
}

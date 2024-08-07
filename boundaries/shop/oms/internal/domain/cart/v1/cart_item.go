package v1

import (
	"github.com/google/uuid"
)

// CartItem represents a cart item.
type CartItem struct {
	// productId is the product ID
	productId uuid.UUID
	// quantity is the quantity of the product
	quantity int32
}

// NewCartItem creates a new CartItem.
func NewCartItem(productId uuid.UUID, quantity int32) CartItem {
	return CartItem{
		productId: productId,
		quantity:  quantity,
	}
}

// GetProductId returns the product ID.
func (c CartItem) GetProductId() uuid.UUID {
	return c.productId
}

// GetQuantity returns the quantity.
func (c CartItem) GetQuantity() int32 {
	return c.quantity
}

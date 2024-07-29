package v1

import (
	"github.com/google/uuid"
)

// CartItem represents a cart item.
type CartItem struct {
	// productId is the product ID
	productId uuid.UUID
	// quantity is the quantity of the product
	quantity int
}

// NewCartItem creates a new CartItem.
func NewCartItem(productId uuid.UUID, quantity int) CartItem {
	return CartItem{
		productId: productId,
		quantity:  quantity,
	}
}

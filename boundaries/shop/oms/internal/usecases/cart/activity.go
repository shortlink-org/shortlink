package cart

import (
	"context"
	"fmt"
)

// AddItemActivity adds an item to the cart
func AddItemActivity(ctx context.Context, item string) (string, error) {
	// Business logic to add an item to the cart
	// For simplicity, we're just returning a success message
	result := fmt.Sprintf("Added item: %s", item)
	return result, nil
}

// RemoveItemActivity removes an item from the cart
func RemoveItemActivity(ctx context.Context, item string) (string, error) {
	// Business logic to remove an item from the cart
	// For simplicity, we're just returning a success message
	result := fmt.Sprintf("Removed item: %s", item)
	return result, nil
}

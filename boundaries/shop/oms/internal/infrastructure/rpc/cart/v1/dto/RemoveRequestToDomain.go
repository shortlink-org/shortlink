package dto

import (
	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

// RemoveRequestToDomain converts a RemoveRequest to a domain model
func RemoveRequestToDomain(r *v1.RemoveRequest) (*domain.CartState, error) {
	// string to uuid
	customerId, err := uuid.Parse(r.CustomerId)
	if err != nil {
		return nil, ErrInvalidCustomerId
	}

	// create a domain model
	state := domain.NewCartState(customerId)

	// remove items from the cart
	for i := range r.GetItems() {
		// string to uuid
		productId, errParseItem := uuid.Parse(r.Items[i].ProductId)
		if errParseItem != nil {
			return nil, ParseItemError{Err: errParseItem, item: r.Items[i].ProductId}
		}

		// create CartItem and remove it from the state
		item := domain.NewCartItem(productId, r.Items[i].Quantity)
		state.AddItem(item)
	}

	return state, nil
}

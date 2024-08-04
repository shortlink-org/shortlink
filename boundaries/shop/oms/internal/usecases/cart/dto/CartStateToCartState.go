package dto

import (
	"github.com/google/uuid"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

// CartStateToCartState converts a CartState to a CartState.
func CartStateToCartState(cartState *v3.CartState) (*v1.CartState, error) {
	customerId, err := uuid.Parse(cartState.GetCustomerId())
	if err != nil {
		return nil, err
	}

	state := v1.NewCartState(customerId)

	for _, item := range cartState.GetItems() {
		productId, err := uuid.Parse(item.GetProductId())
		if err != nil {
			return nil, err
		}

		state.AddItem(v1.NewCartItem(productId, item.GetQuantity()))
	}

	return state, nil
}

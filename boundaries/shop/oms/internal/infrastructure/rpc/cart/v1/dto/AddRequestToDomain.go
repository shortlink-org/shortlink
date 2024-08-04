package dto

import (
	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

// AddRequestToDomain converts an AddRequest to a domain model
func AddRequestToDomain(r *v2.AddRequest) (*domain.CartState, error) {
	// string to uuid
	customerId, err := uuid.Parse(r.CustomerId)
	if err != nil {
		return nil, ErrInvalidCustomerId
	}

	// create a domain model
	item := domain.NewCartState(customerId)

	// add item to the cart
	for i := range r.GetItems() {
		// string to uuid
		productId, errParseItem := uuid.Parse(r.Items[i].ProductId)
		if errParseItem != nil {
			return nil, ParseItemError{Err: errParseItem, item: r.Items[i].ProductId}
		}

		item.AddItem(domain.NewCartItem(productId, r.Items[i].Quantity))
	}

	return item, nil
}

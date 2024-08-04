package v1

import (
	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
)

// AddRequestToDomain converts an AddRequest to a domain model
func AddRequestToDomain(r *AddRequest) (*domain.CartState, error) {
	// string to uuid
	customerId, err := uuid.Parse(r.CustomerId)
	if err != nil {
		return nil, ErrInvalidCustomerId
	}

	// create a domain model
	item := domain.NewCartState(customerId)

	// add item to the cart
	for i := range r.Items {
		// string to uuid
		productId, errParseItem := uuid.Parse(r.Items[i].ProductId)
		if errParseItem != nil {
			return nil, ParseItemError{Err: errParseItem, item: r.Items[i].ProductId}
		}

		item.AddItem(domain.NewCartItem(productId, r.Items[i].Quantity))
	}

	return item, nil
}

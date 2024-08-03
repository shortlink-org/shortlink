package v1

import (
	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
)

// ToDomain converts AddRequest to domain model.
func (r AddRequest) ToDomain() (*domain.CartState, error) {
	// string to uuid
	customerId, err := uuid.Parse(r.CustomerId)
	if err != nil {
		return nil, ErrInvalidCustomerId
	}

	// create a domain model
	item := domain.NewCartState(customerId)

	// add item to the cart
	for id, quantity := range r.Items {
		// string to uuid
		productId, errParseItem := uuid.Parse(id)
		if errParseItem != nil {
			return nil, ParseItemError{Err: errParseItem, item: id}
		}

		item.AddItem(domain.NewCartItem(productId, quantity))
	}

	return &domain.CartState{}, nil
}

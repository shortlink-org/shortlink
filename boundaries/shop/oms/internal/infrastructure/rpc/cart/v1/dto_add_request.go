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
	for _, i := range r.Items {
		// string to uuid
		productId, errParseItem := uuid.Parse(i.ProductId)
		if errParseItem != nil {
			return nil, ParseItemError{Err: errParseItem, item: i.ProductId}
		}

		item.AddItem(domain.NewCartItem(productId, i.Quantity))
	}

	return &domain.CartState{}, nil
}

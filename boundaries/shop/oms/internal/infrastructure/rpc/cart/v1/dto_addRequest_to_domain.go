package v1

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
)

// AddRequestToDomain converts an AddRequest to a domain model
func AddRequestToDomain(r *AddRequest, validator *protovalidate.Validator) (*domain.CartState, error) {
	err := validator.Validate(r)
	if err != nil {
		return nil, err
	}

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

	return item, nil
}

package dto

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
)

// OrderStateToDomain converts a v3.OrderState to a v1.OrderState using the OrderStateBuilder
func OrderStateToDomain(in *v3.OrderState) (*v1.OrderState, error) {
	// Parse the customer ID
	customerID, err := uuid.Parse(in.GetCustomerId())
	if err != nil {
		return nil, fmt.Errorf("invalid customer id: %v", err)
	}

	// Initialize the builder with the customer ID
	builder := v1.NewOrderStateBuilder(customerID)

	// Set the order ID
	orderID, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid order id: %v", err)
	}
	builder.SetId(orderID)

	// Add items to the order
	for _, item := range in.GetItems() {
		productID, err := uuid.Parse(item.GetId())
		if err != nil {
			return nil, fmt.Errorf("invalid item id: %v", err)
		}
		price := decimal.NewFromFloat(item.GetPrice())
		builder.AddItem(productID, item.GetQuantity(), price)
	}

	// Set the status of the order
	builder.SetStatus(context.TODO(), in.GetStatus())

	// Build the OrderState
	orderState, err := builder.Build()
	if err != nil {
		return nil, err
	}

	return orderState, nil
}

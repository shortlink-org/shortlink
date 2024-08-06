package dto

import (
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
)

func DomainToOrderState(in *v1.OrderState) *v2.OrderState {
	// Map Items
	items := make([]*v2.OrderItem, len(in.GetItems()))
	for i, item := range in.GetItems() {
		items[i] = &v2.OrderItem{
			Id:       item.GetProductId().String(),
			Quantity: item.GetQuantity(),
			Price:    item.GetPrice().InexactFloat64(),
		}
	}

	return &v2.OrderState{
		Id:         in.GetOrderID().String(),
		CustomerId: in.GetCustomerId().String(),
		Items:      items,
		Status:     in.GetStatus(),
	}
}

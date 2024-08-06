package dto

import (
	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
)

func OrderStateToDomain(in *v2.OrderState) *v3.OrderState {
	items := make([]*v3.OrderItem, len(in.GetItems()))

	for i, item := range in.GetItems() {
		items[i] = &v3.OrderItem{
			Id:       item.GetProductId().String(),
			Quantity: item.GetQuantity(),
			Price:    item.GetPrice().InexactFloat64(),
		}
	}

	return &v3.OrderState{
		Id:         in.GetOrderID().String(),
		CustomerId: in.GetCustomerId().String(),
		Items:      items,
		Status:     in.GetStatus(),
	}
}

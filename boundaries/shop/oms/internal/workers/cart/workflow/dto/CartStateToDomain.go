package dto

import (
	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

func CartStateToDomain(in *v2.CartState) *v3.CartState {
	items := make([]*v3.CartItem, 0, len(in.GetItems()))

	for _, item := range in.GetItems() {
		items = append(items, &v3.CartItem{
			ProductId: item.GetProductId().String(),
			Quantity:  item.GetQuantity(),
		})
	}

	return &v3.CartState{
		CustomerId: in.GetCustomerId().String(),
		Items:      items,
	}
}

package dto

import (
	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

func GetResponseFromDomain(response *v2.CartState) *v1.GetResponse {
	items := make([]*v1.CartItem, 0, len(response.GetItems()))

	for _, item := range response.GetItems() {
		items = append(items, &v1.CartItem{
			ProductId: item.GetProductId().String(),
			Quantity:  item.GetQuantity(),
		})
	}

	return &v1.GetResponse{
		State: &v1.CartState{
			CustomerId: response.GetCustomerId().String(),
			Items:      items,
		},
	}
}

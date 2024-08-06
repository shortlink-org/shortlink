package v1

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/emptypb"

	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
)

func (o *OrderRPC) Create(ctx context.Context, in *v1.CreateRequest) (*emptypb.Empty, error) {
	// parse order ID to UUID
	orderId, err := uuid.Parse(in.GetOrder().GetId())
	if err != nil {
		return nil, err
	}

	// parse customer ID to UUID
	customerId, err := uuid.Parse(in.GetOrder().GetCustomerId())
	if err != nil {
		return nil, err
	}

	items := make(v2.Items, 0, len(in.GetOrder().GetItems()))
	for _, item := range in.GetOrder().GetItems() {
		// parse product ID to UUID
		productID, err := uuid.Parse(item.GetId())
		if err != nil {
			return nil, err
		}

		price := decimal.NewFromFloat(item.GetPrice())

		items = append(items, v2.NewItem(productID, item.GetQuantity(), price))
	}

	err = o.orderService.Create(ctx, orderId, customerId, items)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

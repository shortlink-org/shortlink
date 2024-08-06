package order

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
)

func (uc *UC) Create(ctx context.Context, customerId uuid.UUID, in v2.Items) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

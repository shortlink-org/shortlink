package order

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (uc *UC) Cancel(ctx context.Context, orderId uuid.UUID) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

package order_rpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (o *Order) OrderHistory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OrderHistoryResponse, error) {
	panic("implement me")
}

func (o *Order) OrderCreate(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*OrderCreateResponse, error) {
	panic("implement me")
}

func (o *Order) OrderUpdate(ctx context.Context, in *OrderUpdateRequest, opts ...grpc.CallOption) (*OrderUpdateResponse, error) {
	panic("implement me")
}

func (o *Order) OrderClose(ctx context.Context, in *OrderCloseRequest, opts ...grpc.CallOption) (*OrderCloseResponse, error) {
	panic("implement me")
}

func (o *Order) OrderApprove(ctx context.Context, in *OrderApproveRequest, opts ...grpc.CallOption) (*OrderApproveResponse, error) {
	panic("implement me")
}

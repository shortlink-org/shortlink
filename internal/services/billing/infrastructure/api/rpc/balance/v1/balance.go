package balance_rpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Balance struct{}

func (b Balance) BalanceHistory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BalanceHistoryResponse, error) {
	panic("implement me")
}

func (b Balance) BalanceUpdate(ctx context.Context, in *BalanceUpdateRequest, opts ...grpc.CallOption) (*BalanceUpdateResponse, error) {
	panic("implement me")
}

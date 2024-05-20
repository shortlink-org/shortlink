package payment_rpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *Payment) Payment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	panic("implement me")
}

func (p *Payment) Payments(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PaymentsResponse, error) {
	return &PaymentsResponse{
		List: nil,
	}, nil
}

func (p *Payment) PaymentCreate(ctx context.Context, in *PaymentCreateRequest, opts ...grpc.CallOption) (*PaymentCreateResponse, error) {
	panic("implement me")
}

func (p *Payment) PaymentClose(ctx context.Context, in *PaymentCloseRequest, opts ...grpc.CallOption) (*PaymentCloseResponse, error) {
	panic("implement me")
}

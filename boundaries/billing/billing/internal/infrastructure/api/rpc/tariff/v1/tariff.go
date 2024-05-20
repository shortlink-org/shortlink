package tariff_rpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *Tariff) Tariff(ctx context.Context, in *TariffRequest, opts ...grpc.CallOption) (*TariffResponse, error) {
	panic("implement me")
}

func (t *Tariff) Tariffs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TariffsResponse, error) {
	panic("implement me")
}

func (t *Tariff) TariffCreate(ctx context.Context, in *TariffCreateRequest, opts ...grpc.CallOption) (*TariffCreateResponse, error) {
	panic("implement me")
}

func (t *Tariff) TariffUpdate(ctx context.Context, in *TariffUpdateRequest, opts ...grpc.CallOption) (*TariffUpdateResponse, error) {
	panic("implement me")
}

func (t *Tariff) TariffClose(ctx context.Context, in *TariffCloseRequest, opts ...grpc.CallOption) (*TariffCloseResponse, error) {
	panic("implement me")
}

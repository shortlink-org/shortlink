package v1

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *CartRPC) Add(ctx context.Context, in *AddRequest) (*emptypb.Empty, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	err = v.Validate(in)
	if err != nil {
		return nil, err
	}

	request, err := in.ToDomain()
	if err != nil {
		return nil, err
	}

	err = c.cartService.Add(ctx, request)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

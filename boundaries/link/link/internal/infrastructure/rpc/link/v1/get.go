package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *LinkRPC) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	_, err := l.service.Get(ctx, in.GetHash())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &GetResponse{
		// Link: resp,
	}, nil
}

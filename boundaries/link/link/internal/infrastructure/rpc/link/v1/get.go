package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *LinkRPC) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	resp, err := l.service.Get(ctx, in.GetHash())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &GetResponse{
		Link: &Link{
			Url:       resp.GetUrl().String(),
			Hash:      resp.GetHash(),
			Describe:  resp.GetDescribe(),
			CreatedAt: resp.GetCreatedAt().GetTimestamp(),
			UpdatedAt: resp.GetUpdatedAt().GetTimestamp(),
		},
	}, nil
}

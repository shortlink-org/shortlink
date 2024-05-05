package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *LinkRPC) Update(ctx context.Context, in *UpdateRequest) (*UpdateResponse, error) {
	if in.GetLink() == nil {
		return nil, ErrEmptyPayload
	}

	entity, err := in.ToEntity()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := l.service.Update(ctx, entity)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &UpdateResponse{
		Link: &Link{
			Url:       resp.GetUrl().String(),
			Hash:      resp.GetHash(),
			Describe:  resp.GetDescribe(),
			CreatedAt: resp.GetCreatedAt().GetTimestamp(),
			UpdatedAt: resp.GetUpdatedAt().GetTimestamp(),
		},
	}, nil
}

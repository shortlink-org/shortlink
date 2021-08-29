package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	queryStore "github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
)

func (l *Link) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	resp, err := l.service.Get(ctx, in.Hash)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &GetResponse{
		Link: resp,
	}, nil
}

func (l *Link) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	// Parse args
	filter := queryStore.Filter{
		Search: &queryStore.StringFilterInput{Contains: &in.Filter},
	}

	resp, err := l.service.List(ctx, &filter)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &ListResponse{
		Links: resp,
	}, nil
}

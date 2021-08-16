package v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	resp, err := l.service.List(ctx, in.Filter)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &ListResponse{
		Links: resp,
	}, nil
}

func (l *Link) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	if in.Link == nil {
		return nil, fmt.Errorf("Create a new link: empty payload")
	}

	resp, err := l.service.Add(ctx, in.Link)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &AddResponse{
		Link: resp,
	}, nil
}

func (l *Link) Update(ctx context.Context, in *UpdateRequest) (*UpdateResponse, error) {
	resp, err := l.service.Update(ctx, in.Link)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &UpdateResponse{
		Link: resp,
	}, nil
}

func (l *Link) Delete(ctx context.Context, in *DeleteRequest) (*empty.Empty, error) {
	_, err := l.service.Delete(ctx, in.Hash)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}

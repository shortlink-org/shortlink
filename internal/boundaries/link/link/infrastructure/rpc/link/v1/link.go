package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
)

func (l *Link) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	resp, err := l.service.Get(ctx, in.GetHash())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &GetResponse{
		Link: resp,
	}, nil
}

func (l *Link) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	// Parse args
	filter := &v1.FilterLink{}

	// TODO: update filter
	// if in.GetFilter() != "" {
	// 	if json.NewDecoder(strings.NewReader(in.GetFilter())).Decode(&filter) != nil {
	// 		return nil, ErrParsePayloadAsString
	// 	}
	// }

	resp, err := l.service.List(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &ListResponse{
		Links: resp,
	}, nil
}

func (l *Link) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	if in.GetLink() == nil {
		return nil, ErrEmptyPayload
	}

	resp, err := l.service.Add(ctx, in.GetLink())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &AddResponse{
		Link: resp,
	}, nil
}

func (l *Link) Update(ctx context.Context, in *UpdateRequest) (*UpdateResponse, error) {
	resp, err := l.service.Update(ctx, in.GetLink())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &UpdateResponse{
		Link: resp,
	}, nil
}

func (l *Link) Delete(ctx context.Context, in *DeleteRequest) (*emptypb.Empty, error) {
	_, err := l.service.Delete(ctx, in.GetHash())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}

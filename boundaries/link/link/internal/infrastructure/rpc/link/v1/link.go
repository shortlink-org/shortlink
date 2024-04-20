package v1

import (
	"context"
	"strings"

	"github.com/segmentio/encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
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

func (l *LinkRPC) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	// Parse args
	filter := &types.FilterLink{}

	if in.GetFilter() != "" {
		if json.NewDecoder(strings.NewReader(in.GetFilter())).Decode(&filter) != nil {
			return nil, ErrParsePayloadAsString
		}
	}

	_, cursor, err := l.service.List(ctx, filter, in.GetCursor(), in.GetLimit())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &ListResponse{
		// Links:  resp,
		Cursor: *cursor,
	}, nil
}

func (l *LinkRPC) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	if in.GetLink() == nil {
		return nil, ErrEmptyPayload
	}

	// _, err := l.service.Add(ctx, in.GetLink())
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }

	return &AddResponse{
		// Link: resp,
	}, nil
}

func (l *LinkRPC) Update(ctx context.Context, in *UpdateRequest) (*UpdateResponse, error) {
	// resp, err := l.service.Update(ctx, in.GetLink())
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }

	return &UpdateResponse{
		// Link: resp,
	}, nil
}

func (l *LinkRPC) Delete(ctx context.Context, in *DeleteRequest) (*emptypb.Empty, error) {
	_, err := l.service.Delete(ctx, in.GetHash())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}

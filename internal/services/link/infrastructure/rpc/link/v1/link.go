package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/segmentio/encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	queryStore "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
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
	// Get session
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("error get metadata from context")
	}

	sess := md.Get("user-id")
	if len(sess) == 0 {
		return nil, errors.New("error get session from metadata")
	}

	// Parse args
	filter := queryStore.Filter{}

	if in.GetFilter() != "" {
		errJsonUnmarshal := json.Unmarshal([]byte(in.GetFilter()), &filter)
		if errJsonUnmarshal != nil {
			return nil, errors.New("error parse payload as string")
		}
	}

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
		return nil, fmt.Errorf("create a new link: empty payload")
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

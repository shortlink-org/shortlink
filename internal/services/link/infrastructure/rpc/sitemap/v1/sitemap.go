package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Sitemap) Parse(ctx context.Context, in *ParseRequest) (*emptypb.Empty, error) {
	err := s.service.Parse(ctx, in.Url)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}

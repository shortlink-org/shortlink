package v1

import (
	"context"
	"strings"

	"github.com/segmentio/encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
)

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

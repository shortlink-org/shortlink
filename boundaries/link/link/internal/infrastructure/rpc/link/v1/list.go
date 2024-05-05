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
		if errDecode := json.NewDecoder(strings.NewReader(in.GetFilter())).Decode(&filter); errDecode != nil {
			return nil, errDecode
		}
	}

	resp, cursor, err := l.service.List(ctx, filter, in.GetCursor(), in.GetLimit())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	links := &Links{
		Link: make([]*Link, 0, len(resp.GetLink())),
	}
	for _, link := range resp.GetLink() {
		links.Link = append(links.Link, &Link{
			Url:       link.GetUrl().String(),
			Hash:      link.GetHash(),
			Describe:  link.GetDescribe(),
			CreatedAt: link.GetCreatedAt().GetTimestamp(),
			UpdatedAt: link.GetUpdatedAt().GetTimestamp(),
		})
	}

	return &ListResponse{
		Links:  links,
		Cursor: *cursor,
	}, nil
}

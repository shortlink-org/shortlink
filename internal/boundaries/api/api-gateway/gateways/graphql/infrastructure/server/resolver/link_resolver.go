package resolver

import (
	"context"
	"fmt"

	"github.com/segmentio/encoding/json"

	link "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/link/v1"
)

// Link ...
func (r *Resolver) Link(ctx context.Context, args struct {
	Hash *string
}) (*LinkResolver, error) {
	response, err := r.LinkServiceClient.Get(ctx, &link_rpc.GetRequest{Hash: *args.Hash})
	if err != nil {
		return nil, err
	}

	return &LinkResolver{
		Link: response.Link,
	}, nil
}

// Links ...
func (r *Resolver) Links(ctx context.Context, args struct {
	Filter *link.FilterLink
}) (*[]*LinkResolver, error) {
	// Filter to string
	filterRaw, err := json.Marshal(args.Filter)
	if err != nil {
		return nil, fmt.Errorf("error parse filter args")
	}

	// Default value for filter; null -> nil
	if string(filterRaw) == "null" {
		filterRaw = nil
	}

	response, err := r.LinkServiceClient.List(ctx, &link_rpc.ListRequest{Filter: string(filterRaw)})
	if err != nil {
		return nil, err
	}

	links := []*LinkResolver{}
	for _, item := range response.Links.Link {
		links = append(links, &LinkResolver{
			Link: item,
		})
	}

	return &links, nil
}

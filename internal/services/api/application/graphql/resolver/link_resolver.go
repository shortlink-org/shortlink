package resolver

import (
	"context"
	"encoding/json"
	"fmt"

	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
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
	Filter *query.Filter
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

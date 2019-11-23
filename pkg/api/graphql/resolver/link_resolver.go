package resolver

import (
	"context"

	"github.com/batazor/shortlink/internal/store/query"
)

// Link ...
func (r *Resolver) Link(ctx context.Context, args struct { //nolint unparam
	Hash *string
}) (*LinkResolver, error) {
	response, err := r.Store.Get(*args.Hash)
	return &LinkResolver{
		Link: response,
	}, err
}

// Links ...
func (r *Resolver) Links(ctx context.Context, args struct {
	Filter *query.Filter
}) (*[]*LinkResolver, error) { // nolint unused
	links := []*LinkResolver{}
	items, err := r.Store.List(args.Filter)
	for _, item := range items {
		links = append(links, &LinkResolver{
			Link: item,
		})
	}

	return &links, err
}

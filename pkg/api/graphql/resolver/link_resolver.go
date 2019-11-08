package resolver

import (
	"context"
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
func (r *Resolver) Links() (*[]*LinkResolver, error) { // nolint unused
	links := []*LinkResolver{}
	items, err := r.Store.List()
	for _, item := range items {
		links = append(links, &LinkResolver{
			Link: item,
		})
	}

	return &links, err
}

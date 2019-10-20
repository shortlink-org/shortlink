package resolver

import (
	"context"
)

func (r *Resolver) Link(ctx context.Context, args struct {
	Hash *string
}) (*LinkResolver, error) {
	response, err := r.Store.Get(*args.Hash)
	return &LinkResolver{
		Link: response,
	}, err
}

func (r *Resolver) Links() (*[]*LinkResolver, error) {
	return &[]*LinkResolver{}, nil
}

package resolver

import (
	"context"
	"github.com/batazor/shortlink/pkg/link"
)

// CreateLink ...
func (r *Resolver) CreateLink(ctx context.Context, args *struct {
	URL      *string
	Hash     *string
	Describe *string
}) (*LinkResolver, error) {
	res, error := r.Store.Add(link.Link{
		URL:      *args.URL,
		Hash:     *args.Hash,
		Describe: *args.Describe,
	})
	return &LinkResolver{
		Link: res,
	}, error
}

// UpdateLink ...
func (*Resolver) UpdateLink(ctx context.Context, args *struct {
	URL      *string
	Hash     *string
	Describe *string
}) (*bool, error) {
	return nil, nil
}

// DeleteLink ...
func (r *Resolver) DeleteLink(ctx context.Context, args *struct {
	Hash *string
}) (*bool, error) {
	error := r.Store.Delete(*args.Hash)
	return nil, error
}

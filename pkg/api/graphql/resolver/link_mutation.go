package resolver

import (
	"context"
	"github.com/batazor/shortlink/pkg/link"
)

func (r *Resolver) CreateLink(ctx context.Context, args *struct {
	Url      *string
	Hash     *string
	Describe *string
}) (*LinkResolver, error) {
	res, error := r.Store.Add(link.Link{
		Url:      *args.Url,
		Hash:     *args.Hash,
		Describe: *args.Describe,
	})
	return &LinkResolver{
		Link: res,
	}, error
}

func (_ *Resolver) UpdateLink(ctx context.Context, args *struct {
	Url      *string
	Hash     *string
	Describe *string
}) (*bool, error) {
	return nil, nil
}

func (r *Resolver) DeleteLink(ctx context.Context, args *struct {
	Hash *string
}) (*bool, error) {
	error := r.Store.Delete(*args.Hash)
	return nil, error
}

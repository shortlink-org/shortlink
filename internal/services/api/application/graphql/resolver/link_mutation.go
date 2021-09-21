package resolver

import (
	"context"

	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
)

// CreateLink ...
func (r *Resolver) CreateLink(ctx context.Context, args *struct { //nolint unused
	URL      *string
	Hash     *string
	Describe *string
}) (*LinkResolver, error) {
	newLink := &v1.Link{
		Url:      *args.URL,
		Hash:     *args.Hash,
		Describe: *args.Describe,
	}

	// Save link
	response, err := r.LinkServiceClient.Add(ctx, &link_rpc.AddRequest{Link: newLink})
	if err != nil {
		return nil, err
	}

	return &LinkResolver{
		Link: response.Link,
	}, err
}

// UpdateLink ...
func (r *Resolver) UpdateLink(ctx context.Context, args *struct { //nolint unused
	URL      *string
	Hash     *string
	Describe *string
}) (*bool, error) {
	updateLink := &v1.Link{
		Url:      *args.URL,
		Hash:     *args.Hash,
		Describe: *args.Describe,
	}

	// Update link
	_, err := r.LinkServiceClient.Update(ctx, &link_rpc.UpdateRequest{Link: updateLink})
	if err != nil {
		return nil, err
	}

	response := true

	return &response, err
}

// DeleteLink ...
func (r *Resolver) DeleteLink(ctx context.Context, args *struct { //nolint unused
	Hash *string
}) (bool, error) {
	_, err := r.LinkServiceClient.Delete(ctx, &link_rpc.DeleteRequest{Hash: *args.Hash})
	if err != nil {
		return false, err
	}

	return true, nil
}

package grpcweb

import (
	"context"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/golang/protobuf/ptypes/empty"
)

// GetLink ...
func (api *API) GetLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	return api.store.Get(req.Hash)
}

// CreateLink ...
func (api *API) CreateLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	return api.store.Add(*req)
}

// DeleteLink ...
func (api *API) DeleteLink(ctx context.Context, req *link.Link) (*empty.Empty, error) {
	response := empty.Empty{}
	return &response, api.store.Delete(req.Hash)
}

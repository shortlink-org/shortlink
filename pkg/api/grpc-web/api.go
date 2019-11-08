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

// GetLinks ...
func (api *API) GetLinks(ctx context.Context, req *link.Link) (*link.Links, error) {
	links, err := api.store.List()
	if err != nil {
		return nil, err
	}

	response := link.Links{}
	for key := range links {
		response.Link = append(response.Link, links[key])
	}

	return &response, nil
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

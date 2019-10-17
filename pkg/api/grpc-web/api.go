package grpc_web

import (
	"context"
	"github.com/batazor/shortlink/pkg/link"
)

func (api *API) GetLink(ctx context.Context, req *GetLinkRequest) (*link.Link, error) {
	return api.store.Get(req.Hash)
}

func (api *API) CreateLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	return api.store.Add(*req)
}

func (api *API) DeleteLink(ctx context.Context, req *DeleteLinkRequest) (*DeleteLinkResponse, error) {
	return nil, api.store.Delete(req.Hash)
}

func (api *API) RedirectLink(ctx context.Context, req *RedirectLinkRequest) (*RedirectLinkResponse, error) {
	_, err := api.store.Get(req.Hash)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

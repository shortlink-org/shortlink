package v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	"github.com/shortlink-org/shortlink/internal/services/api/application/grpc_web/helpers"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// GetLink ...
func (api *API) GetLink(ctx context.Context, req *GetLinkRequest) (*GetLinkResponse, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, v1.METHOD_GET, req.Link.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_GET"})

	// inject spanId in response header
	errSendHeader := helpers.RegisterSpan(ctx)
	if errSendHeader != nil {
		return nil, errSendHeader
	}

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	case notify.Response[any]:
		err := r.Error
		if err != nil {
			return nil, err
		}
		response := r.Payload.(*v1.Link) // nolint:errcheck

		return &GetLinkResponse{
			Link: response,
		}, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")

		return nil, err
	}
}

// GetLinks ...
func (api *API) GetLinks(ctx context.Context, req *GetLinksRequest) (*GetLinksResponse, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, v1.METHOD_LIST, req.Filter, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_LIST"})

	// inject spanId in response header
	errSendHeader := helpers.RegisterSpan(ctx)
	if errSendHeader != nil {
		return nil, errSendHeader
	}

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_LIST")
		return nil, err
	case notify.Response[any]:
		err := r.Error
		if err != nil {
			return nil, err
		}
		links := r.Payload.([]*v1.Link) // nolint:errcheck

		response := v1.Links{}
		response.Link = append(response.Link, links...)

		return &GetLinksResponse{
			Links: response.Link,
		}, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}

// CreateLink ...
func (api *API) CreateLink(ctx context.Context, req *CreateLinkRequest) (*CreateLinkResponse, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, v1.METHOD_ADD, req, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_ADD"})

	// inject spanId in response header
	errSendHeader := helpers.RegisterSpan(ctx)
	if errSendHeader != nil {
		return nil, errSendHeader
	}

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		return nil, err
	case notify.Response[any]:
		err := r.Error
		if err != nil {
			return nil, err
		}
		response := r.Payload.(*v1.Link) // nolint:errcheck

		return &CreateLinkResponse{
			Link: response,
		}, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")

		return nil, err
	}
}

// DeleteLink ...
func (api *API) DeleteLink(ctx context.Context, req *DeleteLinkRequest) (*empty.Empty, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, v1.METHOD_DELETE, req.Link.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_DELETE"})

	// inject spanId in response header
	errSendHeader := helpers.RegisterSpan(ctx)
	if errSendHeader != nil {
		return nil, errSendHeader
	}

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		return &empty.Empty{}, err
	case notify.Response[any]:
		err := r.Error
		if err != nil {
			return nil, err
		}

		return &empty.Empty{}, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		return &empty.Empty{}, err
	}
}

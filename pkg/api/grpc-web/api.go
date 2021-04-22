//go:generate protoc -I. -I../../..  -I../../../third_party/googleapis --go_out=Mpkg/api/grpc-web/grpc-api.proto=.:. --go-grpc_out=Mpkg/api/grpc-web/grpc-api.proto=.:. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative grpc-api.proto
//go:generate protoc -I. -I../../..  -I../../../third_party/googleapis --grpc-gateway_out=logtostderr=true,allow_delete_body=true:. --openapiv2_out=logtostderr=true:. grpc-api.proto
//go:generate mv grpc-api.swagger.json ../../../docs/api.swagger.json

package grpcweb

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain/link"
	"github.com/batazor/shortlink/pkg/api/grpc-web/helpers"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/golang/protobuf/ptypes/empty"
)

// GetLink ...
func (api *API) GetLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, api_type.METHOD_GET, req.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_GET"})

	// inject spanId in response header
	helpers.RegisterSpan(ctx)

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		response := r.Payload.(*link.Link) // nolint errcheck
		return response, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}

// GetLinks ...
func (api *API) GetLinks(ctx context.Context, req *LinkRequest) (*link.Links, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, api_type.METHOD_LIST, req.Filter, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_LIST"})

	// inject spanId in response header
	helpers.RegisterSpan(ctx)

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_LIST")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		links := r.Payload.([]*link.Link) // nolint errcheck

		response := link.Links{}
		for key := range links {
			response.Link = append(response.Link, links[key])
		}

		return &response, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}

// CreateLink ...
func (api *API) CreateLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, api_type.METHOD_ADD, req, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_ADD"})

	// inject spanId in response header
	helpers.RegisterSpan(ctx)

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		response := r.Payload.(*link.Link) // nolint errcheck
		return response, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		return nil, err
	}
}

// DeleteLink ...
func (api *API) DeleteLink(ctx context.Context, req *link.Link) (*empty.Empty, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, api_type.METHOD_DELETE, req.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_DELETE"})

	// inject spanId in response header
	helpers.RegisterSpan(ctx)

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		return &empty.Empty{}, err
	case notify.Response:
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

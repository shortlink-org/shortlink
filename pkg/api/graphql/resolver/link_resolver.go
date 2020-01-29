package resolver

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/internal/store/query"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

// Link ...
func (r *Resolver) Link(ctx context.Context, args struct { //nolint unparam
	Hash *string
}) (*LinkResolver, error) {
	responseCh := make(chan interface{})

	go notify.Publish(api_type.METHOD_GET, *args.Hash, responseCh, "RESPONSE_STORE_GET")

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return &LinkResolver{
				Link: nil,
			}, err
		}
		response := r.Payload.(*link.Link) // nolint errcheck
		return &LinkResolver{
			Link: response,
		}, err
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}

// Links ...
func (r *Resolver) Links(ctx context.Context, args struct { // nolint unused
	Filter *query.Filter
}) (*[]*LinkResolver, error) { // nolint unused
	responseCh := make(chan interface{})

	go notify.Publish(api_type.METHOD_LIST, args.Filter, responseCh, "RESPONSE_STORE_LIST")

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
		responses := r.Payload.([]*link.Link) // nolint errcheck

		links := []*LinkResolver{}
		for _, item := range responses {
			links = append(links, &LinkResolver{
				Link: item,
			})
		}

		return &links, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}

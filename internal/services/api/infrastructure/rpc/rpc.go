/*
RPC Endpoint
*/
package rpc

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

type rpc struct {
	client *grpc.ClientConn

	// Delivery
	MetadataClient metadata_rpc.MetadataClient
	LinkClient     link_rpc.LinkClient
}

func Use(_ context.Context, rpcClient *grpc.ClientConn) (*rpc, error) {
	r := &rpc{
		client: rpcClient,

		// Register clients
		MetadataClient: metadata_rpc.NewMetadataClient(rpcClient),
		LinkClient:     link_rpc.NewLinkClient(rpcClient),
	}

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, r)
	notify.Subscribe(api_type.METHOD_GET, r)
	notify.Subscribe(api_type.METHOD_LIST, r)
	notify.Subscribe(api_type.METHOD_UPDATE, r)
	notify.Subscribe(api_type.METHOD_DELETE, r)

	return r, nil
}

// Notify ...
func (r *rpc) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		{
			resp := notify.Response{
				Name:    "RESPONSE_RPC_ADD",
				Payload: payload,
				Error:   nil,
			}

			g := errgroup.Group{}

			if linkRaw, ok := payload.(*link.Link); ok {
				// Save link
				g.Go(func() error {
					_, err := r.LinkClient.Add(ctx, linkRaw)
					if err != nil {
						return err
					}

					return nil
				})

				// Set Metadata
				g.Go(func() error {
					_, err := r.MetadataClient.Set(ctx, &metadata_rpc.SetMetaRequest{
						Id: linkRaw.Url,
					})
					if err != nil {
						return err
					}

					return nil
				})

				err := g.Wait()
				if err != nil {
					resp.Error = err
				}

				return resp
			}

			resp.Error = errors.New("error parse payload as link.Link")
			return resp
		}
	case api_type.METHOD_GET:
		{
			resp := notify.Response{
				Name:    "RESPONSE_RPC_GET",
				Payload: payload,
				Error:   nil,
			}

			// TODO: use URL address?
			if hash, ok := payload.(string); ok {
				_, err := r.MetadataClient.Get(ctx, &metadata_rpc.GetMetaRequest{
					Id: hash,
				})
				if err != nil {
					resp.Error = err
				}

				return resp
			}

			resp.Error = errors.New("error parse payload as string")
			return resp
		}
	default:
		return notify.Response{}
	}
}

package api_application

import (
	"context"
	"errors"

	"github.com/batazor/shortlink/internal/pkg/notify"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
)

// Notify ...
func (s *Server) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		{
			resp := notify.Response{
				Name:    "RESPONSE_RPC_ADD",
				Payload: payload,
				Error:   nil,
			}

			if linkRaw, ok := payload.(*link.Link); ok {
				// Save link
				_, err := s.LinkClient.Add(ctx, linkRaw)
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
				_, err := s.MetadataClient.Get(ctx, &metadata_rpc.GetMetaRequest{
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

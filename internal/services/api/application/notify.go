package api_application

import (
	"context"
	"errors"

	"github.com/batazor/shortlink/internal/pkg/notify"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc"
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
				linkResp, err := s.LinkClient.Add(ctx, linkRaw)
				if err != nil {
					resp.Error = err
				}

				resp.Payload = linkResp

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
				linkResp, err := s.LinkClient.Get(ctx, &link_rpc.GetRequest{Hash: hash})
				if err != nil {
					resp.Error = err
				}

				resp.Payload = linkResp

				return resp
			}

			resp.Error = errors.New("error parse payload as string")
			return resp
		}
	case api_type.METHOD_LIST:
		{
			resp := notify.Response{
				Name:    "RESPONSE_RPC_LIST",
				Payload: payload,
				Error:   nil,
			}

			// TODO: use URL address?
			if filter, ok := payload.(string); ok {
				_, err := s.LinkClient.List(ctx, &link_rpc.ListRequest{Filter: filter})
				if err != nil {
					resp.Error = err
				}

				return resp
			}

			resp.Error = errors.New("error parse payload as string")
			return resp
		}
	case api_type.METHOD_UPDATE:
		{
			resp := notify.Response{
				Name:    "RESPONSE_RPC_LIST",
				Payload: payload,
				Error:   nil,
			}

			// TODO: use URL address?
			if request, ok := payload.(*link.Link); ok {
				linkResp, err := s.LinkClient.Update(ctx, request)
				if err != nil {
					resp.Error = err
				}

				resp.Payload = linkResp

				return resp
			}

			resp.Error = errors.New("error parse payload as string")
			return resp
		}
	case api_type.METHOD_DELETE:
		{
			resp := notify.Response{
				Name:    "RESPONSE_RPC_LIST",
				Payload: payload,
				Error:   nil,
			}

			// TODO: use URL address?
			if hash, ok := payload.(string); ok {
				linkResp, err := s.LinkClient.Delete(ctx, &link_rpc.DeleteRequest{Hash: hash})
				if err != nil {
					resp.Error = err
				}

				resp.Payload = linkResp

				return resp
			}

			resp.Error = errors.New("error parse payload as string")
			return resp
		}
	default:
		return notify.Response{}
	}
}

package api_application

import (
	"context"
	"errors"

	"github.com/batazor/shortlink/internal/pkg/notify"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
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

			if linkRaw, ok := payload.(*v1.Link); ok {
				// Save link
				linkResp, err := s.LinkCommandServiceClient.Add(ctx, &link_rpc.AddRequest{Link: linkRaw})
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
				linkResp, err := s.LinkQueryServiceClient.Get(ctx, &link_rpc.GetRequest{Hash: hash})
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
				linkResp, err := s.LinkQueryServiceClient.List(ctx, &link_rpc.ListRequest{Filter: filter})
				if err != nil {
					resp.Error = err
				}

				resp.Payload = linkResp

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
			if request, ok := payload.(*v1.Link); ok {
				linkResp, err := s.LinkCommandServiceClient.Update(ctx, &link_rpc.UpdateRequest{Link: request})
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
				Name:    "RESPONSE_RPC_DELETE",
				Payload: payload,
				Error:   nil,
			}

			// TODO: use URL address?
			if hash, ok := payload.(string); ok {
				linkResp, err := s.LinkCommandServiceClient.Delete(ctx, &link_rpc.DeleteRequest{Hash: hash})
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

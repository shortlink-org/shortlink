package api_application

import (
	"context"
	"errors"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	v12 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
)

// Notify ...
func (s *Server) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response { // nolint unused
	switch event {
	case api_domain.METHOD_ADD:
		return s.add(ctx, payload)
	case api_domain.METHOD_GET:
		return s.get(ctx, payload)
	case api_domain.METHOD_LIST:
		return s.list(ctx, payload)
	case api_domain.METHOD_UPDATE:
		return s.update(ctx, payload)
	case api_domain.METHOD_DELETE:
		return s.delete(ctx, payload)
	case api_domain.METHOD_CQRS_GET:
		return s.cqrsGet(ctx, payload)
	default:
		return notify.Response{}
	}
}

func (s *Server) add(ctx context.Context, payload interface{}) notify.Response {
	resp := notify.Response{
		Name:    "RESPONSE_RPC_ADD",
		Payload: payload,
		Error:   nil,
	}

	if linkRaw, ok := payload.(*v1.Link); ok {
		// Save link
		linkResp, err := s.LinkServiceClient.Add(ctx, &link_rpc.AddRequest{Link: linkRaw})
		if err != nil {
			resp.Error = err
		}

		resp.Payload = linkResp

		return resp
	}

	resp.Error = errors.New("error parse payload as link.Link")
	return resp
}

func (s *Server) get(ctx context.Context, payload interface{}) notify.Response {
	resp := notify.Response{
		Name:    "RESPONSE_RPC_GET",
		Payload: payload,
		Error:   nil,
	}

	// TODO: use URL address?
	if hash, ok := payload.(string); ok {
		linkResp, err := s.LinkServiceClient.Get(ctx, &link_rpc.GetRequest{Hash: hash})
		if err != nil {
			resp.Error = err
		}

		resp.Payload = linkResp.Link

		return resp
	}

	resp.Error = errors.New("error parse payload as string")
	return resp
}

func (s *Server) list(ctx context.Context, payload interface{}) notify.Response {
	resp := notify.Response{
		Name:    "RESPONSE_RPC_LIST",
		Payload: payload,
		Error:   nil,
	}

	// TODO: use URL address?
	if filter, ok := payload.(string); ok {
		linkResp, err := s.LinkServiceClient.List(ctx, &link_rpc.ListRequest{Filter: filter})
		if err != nil {
			resp.Error = err
		}

		resp.Payload = linkResp

		return resp
	}

	resp.Error = errors.New("error parse payload as string")
	return resp
}

func (s *Server) update(ctx context.Context, payload interface{}) notify.Response {
	resp := notify.Response{
		Name:    "RESPONSE_RPC_LIST",
		Payload: payload,
		Error:   nil,
	}

	// TODO: use URL address?
	if request, ok := payload.(*v1.Link); ok {
		linkResp, err := s.LinkServiceClient.Update(ctx, &link_rpc.UpdateRequest{Link: request})
		if err != nil {
			resp.Error = err
		}

		resp.Payload = linkResp

		return resp
	}

	resp.Error = errors.New("error parse payload as string")
	return resp
}

func (s *Server) delete(ctx context.Context, payload interface{}) notify.Response {
	resp := notify.Response{
		Name:    "RESPONSE_RPC_DELETE",
		Payload: payload,
		Error:   nil,
	}

	// TODO: use URL address?
	if hash, ok := payload.(string); ok {
		linkResp, err := s.LinkServiceClient.Delete(ctx, &link_rpc.DeleteRequest{Hash: hash})
		if err != nil {
			resp.Error = err
		}

		resp.Payload = linkResp

		return resp
	}

	resp.Error = errors.New("error parse payload as string")
	return resp
}

func (s *Server) cqrsGet(ctx context.Context, payload interface{}) notify.Response {
	resp := notify.Response{
		Name:    "RESPONSE_RPC_CQRS_GET",
		Payload: payload,
		Error:   nil,
	}

	// TODO: use URL address?
	if hash, ok := payload.(string); ok {
		linkResp, err := s.LinkQueryServiceClient.Get(ctx, &v12.GetRequest{Hash: hash})
		if err != nil {
			resp.Error = err
		}

		if linkResp != nil {
			resp.Payload = linkResp.Link
		}

		return resp
	}

	resp.Error = errors.New("error parse payload as string")
	return resp
}

//go:generate protoc -I. -I../../domain/link --go_out=Minternal/link/domain/link.proto=.:. --go-grpc_out=Minternal/link/domain/link.proto=.:. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative link_rpc.proto

/*
Link Service. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package link_rpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/application"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	link_store "github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/rpc"
)

type Link struct {
	UnimplementedLinkServer

	service *application.Service
	log     logger.Logger
}

func New(runRPCServer *rpc.RPCServer, st *link_store.LinkStore, log logger.Logger) (*Link, error) {
	server := &Link{
		// Create Service Application
		service: &application.Service{
			Store: st,
		},
		log: log,
	}

	// Register services
	RegisterLinkServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *Link) Add(ctx context.Context, in *link.Link) (*link.Link, error) {
	var err error
	responseCh := make(chan interface{})

	// TODO: send []byte format
	go notify.Publish(ctx, api_type.METHOD_ADD, in, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_ADD"})

	c := <-responseCh
	switch resp := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		return nil, err
	case notify.Response:
		if resp.Error != nil {
			return nil, resp.Error
		}

		return resp.Payload.(*link.Link), nil
	default:
		return nil, status.Error(codes.InvalidArgument, "default case")
	}
}

func (m *Link) Get(ctx context.Context, in *link.Link) (*link.Link, error) {
	panic("implement me")
}

func (m *Link) List(ctx context.Context, in *LinkRequest) (*link.Links, error) {
	panic("implement me")
}

func (m *Link) Update(ctx context.Context, in *link.Link) (*link.Link, error) {
	panic("implement me")
}

func (m *Link) Delete(ctx context.Context, in *link.Link) (*link.Link, error) {
	panic("implement me")
}

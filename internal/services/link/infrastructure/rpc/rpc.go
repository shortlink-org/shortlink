//go:generate protoc -I. -I../../domain/link --go_out=Minternal/link/domain/link.proto=.:. --go-grpc_out=Minternal/link/domain/link.proto=.:. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative link_rpc.proto

/*
Link Service. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package link_rpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/batazor/shortlink/internal/pkg/logger"
	link_application "github.com/batazor/shortlink/internal/services/link/application"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	"github.com/batazor/shortlink/pkg/rpc"
)

type Link struct {
	UnimplementedLinkServer

	service *link_application.Service
	log     logger.Logger
}

func New(runRPCServer *rpc.RPCServer, application *link_application.Service, log logger.Logger) (*Link, error) {
	server := &Link{
		// Create Service Application
		service: application,
		log:     log,
	}

	// Register services
	RegisterLinkServer(runRPCServer.Server, server)
	go runRPCServer.Run()

	return server, nil
}

func (l *Link) Add(ctx context.Context, in *link.Link) (*link.Link, error) {
	resp, err := l.service.AddLink(ctx, in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (l *Link) Get(ctx context.Context, in *GetRequest) (*link.Link, error) {
	resp, err := l.service.GetLink(ctx, in.Hash)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (l *Link) List(ctx context.Context, in *ListRequest) (*link.Links, error) {
	resp, err := l.service.ListLink(ctx, in.Filter)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (l *Link) Update(ctx context.Context, in *link.Link) (*link.Link, error) {
	resp, err := l.service.UpdateLink(ctx, in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (l *Link) Delete(ctx context.Context, in *DeleteRequest) (*empty.Empty, error) {
	_, err := l.service.DeleteLink(ctx, in.Hash)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}

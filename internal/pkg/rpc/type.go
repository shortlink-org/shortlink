package rpc

import "google.golang.org/grpc"

type RPCServer struct {
	Run      func()
	Server   *grpc.Server
	Endpoint string
}

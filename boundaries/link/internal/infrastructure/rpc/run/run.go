package run

import (
	rpc "github.com/shortlink-org/go-sdk/grpc"
)

type Response struct{}

func Run(runRPCServer *rpc.Server) (*Response, error) {
	if runRPCServer != nil {
		go runRPCServer.Run()
	}

	return nil, nil
}

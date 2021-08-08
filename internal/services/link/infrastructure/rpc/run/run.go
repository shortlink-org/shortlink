package run

import (
	"github.com/batazor/shortlink/pkg/rpc"
)

type Response struct{}

func Run(runRPCServer *rpc.RPCServer) (*Response, error) {
	go runRPCServer.Run()
	return nil, nil
}

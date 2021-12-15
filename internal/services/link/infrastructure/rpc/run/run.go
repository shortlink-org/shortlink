package run

import (
	"github.com/batazor/shortlink/pkg/rpc"
)

type Response struct{}

func Run(runRPCServer *rpc.RPCServer) (*Response, error) {
	if runRPCServer != nil {
		go runRPCServer.Run()
	}

	return nil, nil
}

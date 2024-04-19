package run

import (
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type Response struct{}

func Run(runRPCServer *rpc.Server) (*Response, error) {
	if runRPCServer != nil {
		go runRPCServer.Run()
	}

	return nil, nil
}

/*
Cart UC. Infrastructure layer. RPC Endpoint
*/

package v1

import (
	"go.temporal.io/sdk/client"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type CartRPC struct {
	CartServiceServer

	log logger.Logger

	client client.Client
}

func New(runRPCServer *rpc.Server, log logger.Logger, c client.Client) (*CartRPC, error) {
	server := &CartRPC{
		log:    log,
		client: c,
	}

	// Register services
	if runRPCServer != nil {
		RegisterCartServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}

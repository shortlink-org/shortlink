/*
Cart UC. Infrastructure layer. RPC Endpoint
*/

package v1

import (
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/cart"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type CartRPC struct {
	CartServiceServer

	log logger.Logger

	cartService *cart.UC
}

func New(runRPCServer *rpc.Server, log logger.Logger, cartService *cart.UC) (*CartRPC, error) {
	server := &CartRPC{
		log:         log,
		cartService: cartService,
	}

	// Register services
	if runRPCServer != nil {
		RegisterCartServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}

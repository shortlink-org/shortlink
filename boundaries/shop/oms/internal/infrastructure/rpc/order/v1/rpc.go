/*
Cart UC. Infrastructure layer. RPC Endpoint
*/

package v1

import (
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/cart"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type OrderRPC struct {
	OrderServiceServer

	// Common
	log logger.Logger

	// Services
	cartService *cart.UC
}

func New(runRPCServer *rpc.Server, log logger.Logger, cartService *cart.UC) (*OrderRPC, error) {
	server := &OrderRPC{
		// Common
		log: log,

		// Services
		cartService: cartService,
	}

	// Register services
	if runRPCServer != nil {
		RegisterOrderServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}

/*
Cart UC. Infrastructure layer. RPC Endpoint
*/

package v1

import (
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/order"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type OrderRPC struct {
	OrderServiceServer

	// Common
	log logger.Logger

	// Services
	orderService *order.UC
}

func New(runRPCServer *rpc.Server, log logger.Logger, orderService *order.UC) (*OrderRPC, error) {
	server := &OrderRPC{
		// Common
		log: log,

		// Services
		orderService: orderService,
	}

	// Register services
	if runRPCServer != nil {
		RegisterOrderServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}

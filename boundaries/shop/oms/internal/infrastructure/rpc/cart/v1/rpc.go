/*
Cart UC. Infrastructure layer. RPC Endpoint
*/

package v1

import (
	"github.com/bufbuild/protovalidate-go"

	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/cart"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type CartRPC struct {
	CartServiceServer

	// Common
	log       logger.Logger
	validator *protovalidate.Validator

	// Services
	cartService *cart.UC
}

func New(runRPCServer *rpc.Server, log logger.Logger, cartService *cart.UC) (*CartRPC, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	server := &CartRPC{
		// Common
		log:       log,
		validator: validator,

		// Services
		cartService: cartService,
	}

	// Register services
	if runRPCServer != nil {
		RegisterCartServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}

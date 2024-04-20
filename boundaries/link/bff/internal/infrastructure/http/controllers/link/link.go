package link

import (
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

type Controller struct {
	log logger.Logger

	linkServiceClient link_rpc.LinkServiceClient
}

// NewController - create new link controller
func NewController(log logger.Logger, linkServiceClient link_rpc.LinkServiceClient) Controller {
	return Controller{
		log: log,

		linkServiceClient: linkServiceClient,
	}
}

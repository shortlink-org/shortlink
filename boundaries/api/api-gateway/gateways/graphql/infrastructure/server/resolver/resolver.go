package resolver

import (
	"github.com/shortlink-org/shortlink/pkg/db"

	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
)

// Resolver ...
type Resolver struct {
	Store             db.DB
	LinkServiceClient link_rpc.LinkServiceClient
}

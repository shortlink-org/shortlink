package resolver

import (
	link_rpc "github.com/shortlink-org/shortlink/boundaries/link/link/infrastructure/rpc/link/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
)

// Resolver ...
type Resolver struct {
	Store             db.DB
	LinkServiceClient link_rpc.LinkServiceClient
}

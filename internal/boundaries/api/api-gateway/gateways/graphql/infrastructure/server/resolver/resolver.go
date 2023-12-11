package resolver

import (
	link_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/link/v1"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
)

// Resolver ...
type Resolver struct {
	Store             db.DB
	LinkServiceClient link_rpc.LinkServiceClient
}

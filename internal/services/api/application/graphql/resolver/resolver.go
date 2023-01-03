package resolver

import (
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	link_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
)

// Resolver ...
type Resolver struct {
	Store             db.DB
	LinkServiceClient link_rpc.LinkServiceClient
}

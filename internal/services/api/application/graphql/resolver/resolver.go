package resolver

import (
	"github.com/batazor/shortlink/internal/pkg/db"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
)

// Resolver ...
type Resolver struct { //nolint unused
	Store             db.DB
	LinkServiceClient link_rpc.LinkServiceClient
}

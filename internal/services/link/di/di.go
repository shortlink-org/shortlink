//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/batazor/shortlink/internal/pkg/logger"
	link_store "github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
)

type LinkService struct {
	Logger logger.Logger

	// Delivery
	metadataRPCServer *metadata_rpc.MetadataServer

	// Repository
	LinkStore *link_store.LinkStore
}

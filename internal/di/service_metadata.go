//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/logger"
	meta_store "github.com/batazor/shortlink/internal/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

// MetadataService =====================================================================================================
var MetadataSet = wire.NewSet(
	DefaultSet,
	NewMetadataService,
	InitStore,
	rpc.RunGRPCServer,
	InitMetaStore,
	InitSentry,
	InitMonitoring,
)

func NewMetadataService(
	log logger.Logger,
	autoMaxProcsOption diAutoMaxPro,
	db *db.Store,
	serverRPC *rpc.RPCServer,
	metaStore *meta_store.MetaStore,
	monitoring *http.ServeMux,
	sentryHandler *sentryhttp.Handler,
) (*Service, error) {
	return &Service{
		Log:        log,
		ServerRPC:  serverRPC,
		DB:         db,
		MetaStore:  metaStore,
		Monitoring: monitoring,
		Sentry:     sentryHandler,
	}, nil
}

func InitializeMetadataService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(MetadataSet))
}

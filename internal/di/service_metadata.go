//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	meta_di "github.com/batazor/shortlink/internal/services/metadata/di"
	"github.com/batazor/shortlink/pkg/rpc"
)

type ServiceMetadata struct {
	Service

	MetaService *meta_di.MetaDataService
}

// InitMetaService =====================================================================================================
func InitMetaService(ctx context.Context, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store) (*meta_di.MetaDataService, func(), error) {
	return meta_di.InitializeMetaDataService(ctx, runRPCServer, log, db)
}

// MetadataService =====================================================================================================
var MetadataSet = wire.NewSet(
	DefaultSet,
	store.New,
	rpc.InitServer,
	sentry.New,
	monitoring.New,
	InitMetaService,
	NewMetadataService,
)

func NewMetadataService(
	log logger.Logger,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	db *db.Store,
	serverRPC *rpc.RPCServer,
	monitoring *http.ServeMux,
	sentryHandler *sentryhttp.Handler,
	metadataService *meta_di.MetaDataService,
) (*ServiceMetadata, error) {
	return &ServiceMetadata{
		Service: Service{
			Log:        log,
			ServerRPC:  serverRPC,
			DB:         db,
			Monitoring: monitoring,
			Sentry:     sentryHandler,
		},
		MetaService: metadataService,
	}, nil
}

func InitializeMetadataService() (*ServiceMetadata, func(), error) {
	panic(wire.Build(MetadataSet))
}

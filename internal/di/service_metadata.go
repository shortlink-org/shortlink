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
	meta_store "github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

type ServiceMetadata struct {
	Service
	MetaStore *meta_store.MetaStore
}

// InitMetaStore =======================================================================================================
func InitMetaStore(ctx context.Context, log logger.Logger, conn *db.Store) (*meta_store.MetaStore, error) {
	st := meta_store.MetaStore{}
	metaStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return metaStore, nil
}

// MetadataService =====================================================================================================
var MetadataSet = wire.NewSet(
	DefaultSet,
	store.New,
	rpc.InitServer,
	InitMetaStore,
	sentry.New,
	monitoring.New,
	NewMetadataService,
)

func NewMetadataService(
	log logger.Logger,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	db *db.Store,
	serverRPC *rpc.RPCServer,
	metaStore *meta_store.MetaStore,
	monitoring *http.ServeMux,
	sentryHandler *sentryhttp.Handler,
) (*ServiceMetadata, error) {
	return &ServiceMetadata{
		Service: Service{
			Log:        log,
			ServerRPC:  serverRPC,
			DB:         db,
			Monitoring: monitoring,
			Sentry:     sentryHandler,
		},
		MetaStore: metaStore,
	}, nil
}

func InitializeMetadataService() (*ServiceMetadata, func(), error) {
	panic(wire.Build(MetadataSet))
}

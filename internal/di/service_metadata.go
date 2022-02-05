//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/profiling"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	meta_di "github.com/batazor/shortlink/internal/services/metadata/di"
	"github.com/batazor/shortlink/pkg/rpc"
)

type ServiceMetadata struct {
	Service

	MetaService *meta_di.MetaDataService
}

// InitMetaService =====================================================================================================
func InitMetaDataService(ctx context.Context, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq v1.MQ) (*meta_di.MetaDataService, func(), error) {
	return meta_di.InitializeMetaDataService(ctx, runRPCServer, log, db, mq)
}

// MetadataService =====================================================================================================
var MetadataSet = wire.NewSet(
	DefaultSet,
	store.New,
	rpc.InitServer,
	sentry.New,
	monitoring.New,
	profiling.New,
	mq_di.New,
	InitMetaDataService,
	NewMetadataService,
)

func NewMetadataService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	mq v1.MQ,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	db *db.Store,
	serverRPC *rpc.RPCServer,
	monitoring *http.ServeMux,
	sentryHandler *sentryhttp.Handler,
	metadataService *meta_di.MetaDataService,
) (*ServiceMetadata, error) {
	return &ServiceMetadata{
		Service: Service{
			Ctx:        ctx,
			Log:        log,
			ServerRPC:  serverRPC,
			DB:         db,
			MQ:         mq,
			Monitoring: monitoring,
			Sentry:     sentryHandler,
		},
		MetaService: metadataService,
	}, nil
}

func InitializeMetadataService() (*ServiceMetadata, func(), error) {
	panic(wire.Build(MetadataSet))
}

//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"google.golang.org/grpc"

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
	link_di "github.com/batazor/shortlink/internal/services/link/di"
	"github.com/batazor/shortlink/pkg/rpc"
)

type ServiceLink struct {
	Service

	LinkService *link_di.LinkService
}

// InitLinkService =====================================================================================================
func InitLinkService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq v1.MQ, cache *cache.Cache) (*link_di.LinkService, func(), error) {
	return link_di.InitializeLinkService(ctx, runRPCClient, runRPCServer, log, db, mq, cache)
}

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(
	DefaultSet,
	store.New,
	sentry.New,
	monitoring.New,
	profiling.New,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
	InitLinkService,
	NewLinkService,
)

func NewLinkService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	mq v1.MQ,
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	db *db.Store,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
	linkService *link_di.LinkService,
) (*ServiceLink, error) {
	return &ServiceLink{
		Service: Service{
			Ctx:           ctx,
			Log:           log,
			MQ:            mq,
			Tracer:        tracer,
			Monitoring:    monitoring,
			Sentry:        sentryHandler,
			DB:            db,
			PprofEndpoint: pprofHTTP,
			ClientRPC:     clientRPC,
			ServerRPC:     serverRPC,
		},
		LinkService: linkService,
	}, nil
}

func InitializeLinkService() (*ServiceLink, func(), error) {
	panic(wire.Build(LinkSet))
}

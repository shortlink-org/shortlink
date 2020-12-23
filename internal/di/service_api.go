//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	link_store "github.com/batazor/shortlink/internal/api/infrastructure/store"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/profiling"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/pkg/rpc"
)

// InitLinkStore =======================================================================================================
func InitLinkStore(ctx context.Context, log logger.Logger, conn *db.Store) (*link_store.LinkStore, error) {
	st := link_store.LinkStore{}
	linkStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	DefaultSet,
	store.New,
	InitLinkStore,
	sentry.New,
	monitoring.New,
	profiling.New,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
	NewAPIService,
)

func NewAPIService(
	ctx context.Context,
	log logger.Logger,
	mq mq.MQ,
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	db *db.Store,
	linkStore *link_store.LinkStore,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Ctx:           ctx,
		Log:           log,
		MQ:            mq,
		Tracer:        tracer,
		Monitoring:    monitoring,
		Sentry:        sentryHandler,
		DB:            db,
		LinkStore:     linkStore,
		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

func InitializeAPIService() (*Service, func(), error) {
	panic(wire.Build(APISet))
}

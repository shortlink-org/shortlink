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

	api_mq "github.com/batazor/shortlink/internal/services/link/infrastructure/mq"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/profiling"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/internal/pkg/notify"
	link_store "github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/rpc"
)

type ServiceLink struct {
	Service
	LinkMQ    *api_mq.Event
	LinkStore *link_store.LinkStore
}

// InitLinkMQ ==========================================================================================================
func InitLinkMQ(ctx context.Context, log logger.Logger, mq mq.MQ) (*api_mq.Event, error) {
	linkMQ := &api_mq.Event{
		MQ: mq,
	}

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, linkMQ)

	return linkMQ, nil
}

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
var LinkSet = wire.NewSet(
	DefaultSet,
	store.New,
	InitLinkMQ,
	InitLinkStore,
	sentry.New,
	monitoring.New,
	profiling.New,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
	NewLinkService,
)

func NewLinkService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	mq mq.MQ,
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	db *db.Store,
	api_mq *api_mq.Event,
	linkStore *link_store.LinkStore,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
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
		LinkMQ:    api_mq,
		LinkStore: linkStore,
	}, nil
}

func InitializeLinkService() (*ServiceLink, func(), error) {
	panic(wire.Build(LinkSet))
}

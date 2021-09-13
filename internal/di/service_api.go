//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/text/message"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	"github.com/batazor/shortlink/internal/di/internal/profiling"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/pkg/logger"
	api_di "github.com/batazor/shortlink/internal/services/api/di"
	"github.com/batazor/shortlink/pkg/rpc"
)

type ServiceAPI struct {
	Service

	APIService *api_di.APIService
}

// InitAPIService =====================================================================================================
func InitAPIService(ctx context.Context, i18n *message.Printer, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, tracer *opentracing.Tracer) (*api_di.APIService, func(), error) {
	return api_di.InitializeAPIService(ctx, i18n, runRPCClient, runRPCServer, log, tracer)
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	DefaultSet,
	sentry.New,
	monitoring.New,
	profiling.New,
	rpc.InitServer,
	rpc.InitClient,
	InitAPIService,
	NewAPIService,
)

func NewAPIService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	i18n *message.Printer,
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	clientRPC *grpc.ClientConn,
	apiService *api_di.APIService,
) (*ServiceAPI, error) {
	return &ServiceAPI{
		Service: Service{
			Ctx:           ctx,
			Log:           log,
			Tracer:        tracer,
			I18N:          i18n,
			Monitoring:    monitoring,
			Sentry:        sentryHandler,
			PprofEndpoint: pprofHTTP,
			ClientRPC:     clientRPC,
		},
		APIService: apiService,
	}, nil
}

func InitializeAPIService() (*ServiceAPI, func(), error) {
	panic(wire.Build(APISet))
}

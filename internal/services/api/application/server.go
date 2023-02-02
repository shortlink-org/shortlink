package api_application

import (
	"context"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/message"

	http_server "github.com/shortlink-org/shortlink/pkg/http/server"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/services/api/application/cloudevents"
	"github.com/shortlink-org/shortlink/internal/services/api/application/graphql"
	grpcweb "github.com/shortlink-org/shortlink/internal/services/api/application/grpc_web/v1"
	http_chi "github.com/shortlink-org/shortlink/internal/services/api/application/http-chi"
	link_cqrs "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

// runAPIServer - start HTTP-server
func RunAPIServer(
	ctx context.Context,
	i18n *message.Printer,
	log logger.Logger,
	rpcServer *rpc.RPCServer,
	tracer *trace.TracerProvider,
	monitoring *http.ServeMux,

	// delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) (*API, error) {
	var server API

	viper.SetDefault("API_TYPE", "http-chi") // Select: http-chi, gRPC-web, graphql, cloudevents
	// API port
	viper.SetDefault("API_PORT", 7070) // nolint:gomnd
	// Request Timeout (seconds)
	viper.SetDefault("API_TIMEOUT", 60) // nolint:gomnd

	config := http_server.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT") * time.Second, // nolint:durationcheck
	}

	serverType := viper.GetString("API_TYPE")

	switch serverType {
	case "http-chi":
		server = &http_chi.API{}
	case "gRPC-web":
		server = &grpcweb.API{
			RPC: rpcServer,
		}
	case "graphql":
		server = &graphql.API{}
	case "cloudevents":
		server = &cloudevents.API{}
	default:
		server = &http_chi.API{}
	}

	g := errgroup.Group{}

	g.Go(func() error {
		return server.Run(ctx, i18n, config, log, tracer, link_rpc, link_command, link_query, sitemap_rpc)
	})

	return &server, nil
}

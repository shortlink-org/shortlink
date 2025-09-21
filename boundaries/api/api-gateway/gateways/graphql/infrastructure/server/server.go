package server

import (
	"context"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"

	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

// RunAPIServer - start HTTP-server
func RunAPIServer(
	ctx context.Context,
	log logger.Logger,
	rpcServer *rpc.Server,
	tracer trace.TracerProvider,
	monitor *metrics.Monitoring,

	// delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) (*API, error) {
	viper.SetDefault("BASE_PATH", "/api") // Base path for API endpoints
	// API port
	viper.SetDefault("API_PORT", 7070) //nolint:mnd
	// Request Timeout (seconds)
	viper.SetDefault("API_TIMEOUT", "60s")

	config := http_server.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	server := &API{}

	g := errgroup.Group{}

	g.Go(func() error {
		return server.Run(ctx, config, log, tracer, link_rpc, link_command, link_query, sitemap_rpc)
	})

	return server, nil
}

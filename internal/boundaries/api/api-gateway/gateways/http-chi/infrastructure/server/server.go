package server

import (
	"context"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/message"

	http_server "github.com/shortlink-org/shortlink/internal/pkg/http/server"
	_ "github.com/shortlink-org/shortlink/internal/pkg/i18n"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/rpc"
	link_cqrs "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
)

// RunAPIServer - start HTTP-server
func RunAPIServer(
	ctx context.Context,
	i18n *message.Printer,
	log logger.Logger,
	rpcServer *rpc.Server,
	tracer trace.TracerProvider,
	monitor *monitoring.Monitoring,

	// delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) (*API, error) {
	viper.SetDefault("BASE_PATH", "/api") // Base path for API endpoints
	// API port
	viper.SetDefault("API_PORT", 7070) // nolint:gomnd
	// Request Timeout (seconds)
	viper.SetDefault("API_TIMEOUT", "60s")

	config := http_server.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	server := &API{}

	g := errgroup.Group{}

	g.Go(func() error {
		return server.Run(ctx, i18n, config, log, tracer, monitor, link_rpc, link_command, link_query, sitemap_rpc)
	})

	return server, nil
}

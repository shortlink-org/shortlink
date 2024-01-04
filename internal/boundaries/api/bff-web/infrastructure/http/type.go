package http

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/internal/boundaries/api/bff-web/infrastructure/http/controllers/cqrs"
	"github.com/shortlink-org/shortlink/internal/boundaries/api/bff-web/infrastructure/http/controllers/link"
	"github.com/shortlink-org/shortlink/internal/boundaries/api/bff-web/infrastructure/http/controllers/sitemap"
	link_cqrs "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/sitemap/v1"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/rpc"
)

// API HTTP-server
type Server struct {
	ctx    context.Context
	jsonpb protojson.MarshalOptions
}

// Config HTTP-server config
type Config struct {
	// Common
	Ctx    context.Context
	I18n   *message.Printer
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	RpcServer *rpc.Server

	// Infrastructure
	Link_rpc     link_rpc.LinkServiceClient
	Link_command link_cqrs.LinkCommandServiceClient
	Link_query   link_cqrs.LinkQueryServiceClient
	Sitemap_rpc  sitemap_rpc.SitemapServiceClient
}

type Controller struct {
	link.Controller
	cqrs.LinkCQRSController
	sitemap.SitemapController
}

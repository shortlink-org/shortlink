/*
API
*/

package api_application

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/pkg/logger"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
)

// API - general describe of API
type API interface { // nolint unused
	Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer *opentracing.Tracer) error
}

type Server struct {
	// Delivery
	MetadataClient           metadata_rpc.MetadataClient
	LinkCommandServiceClient link_rpc.LinkCommandServiceClient
	LinkQueryServiceClient   link_rpc.LinkQueryServiceClient
}

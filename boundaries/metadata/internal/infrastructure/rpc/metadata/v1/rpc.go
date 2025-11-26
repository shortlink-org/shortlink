// Package v1 exposes the RPC transport layer for metadata use cases.
package v1

import (
	"context"

	rpc "github.com/shortlink-org/go-sdk/grpc"
	"github.com/shortlink-org/go-sdk/logger"

	infraerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/errors"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/metadata"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/parsers"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/screenshot"
)

type Metadata struct {
	MetadataServiceServer

	// application
	parserUC     *parsers.UC
	screenshotUC *screenshot.UC
	metadataUC   *metadata.UC

	// common
	log logger.Logger
}

//nolint:revive // wiring requires passing multiple collaborators
func New(
	log logger.Logger,
	runRPCServer *rpc.Server,
	parsersUC *parsers.UC,
	screenshotUC *screenshot.UC,
	metadataUC *metadata.UC,
) (
	*Metadata,
	error,
) {
	server := &Metadata{
		// application
		parserUC:     parsersUC,
		screenshotUC: screenshotUC,
		metadataUC:   metadataUC,

		// common
		log: log,
	}

	// Register services
	if runRPCServer != nil {
		RegisterMetadataServiceServer(runRPCServer.Server, server)

		go runRPCServer.Run()
	}

	return server, nil
}

func (m *Metadata) Get(ctx context.Context, req *MetadataServiceGetRequest) (*MetadataServiceGetResponse, error) {
	// Get metadata
	meta, err := m.parserUC.Get(ctx, req.GetUrl())
	if err != nil {
		dto := infraerrors.FromDomainError("metadata.rpc", err)
		return nil, dto.ToGRPC()
	}

	// Get screenshotURL
	// screenshotURL, err := m.screenshotUC.Get(ctx, fmt.Sprintf("%s.png", req.GetUrl()))
	// if err != nil {
	// 	return nil, err
	// }

	// Set screenshotURL
	// meta.ImageUrl = screenshotURL.String()

	return &MetadataServiceGetResponse{
		Meta: meta,
	}, nil
}

func (m *Metadata) Set(ctx context.Context, req *MetadataServiceSetRequest) (*MetadataServiceSetResponse, error) {
	// Set metadata
	meta, err := m.metadataUC.Add(ctx, req.GetUrl())
	if err != nil {
		dto := infraerrors.FromDomainError("metadata.rpc", err)
		return nil, dto.ToGRPC()
	}

	return &MetadataServiceSetResponse{
		Meta: meta,
	}, nil
}

package v1

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	repository_err "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/error"
)

// mapDomainErrorToGRPC maps domain errors to appropriate gRPC status codes
// Transport layer works only with DomainError interface, not with concrete error types
// Domain owns the semantics of error mapping via GRPCCode() method
func mapDomainErrorToGRPC(err error) error {
	if err == nil {
		return nil
	}

	// Check if error implements DomainError interface
	var de domain.DomainError
	if errors.As(err, &de) {
		// Domain error knows its own gRPC code
		return status.Error(de.GRPCCode(), err.Error())
	}

	// Repository errors
	if errors.Is(err, repository_err.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}

	// Fallback for unknown errors
	return status.Error(codes.Internal, err.Error())
}

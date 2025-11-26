package v1

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	repository_err "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/error"
)

// mapDomainErrorToGRPC maps domain errors to appropriate gRPC status codes.
// Transport layer works with explicit domain error types to keep domain independent from transport concerns.
func mapDomainErrorToGRPC(err error) error {
	if err == nil {
		return nil
	}

	var linkErr *domain.LinkError
	if errors.As(err, &linkErr) {
		switch linkErr.Code() {
		case domain.CodeNotFound:
			return status.Error(codes.NotFound, linkErr.Error())
		case domain.CodeInvalidInput:
			return status.Error(codes.InvalidArgument, linkErr.Error())
		case domain.CodePermissionDenied:
			return status.Error(codes.PermissionDenied, linkErr.Error())
		case domain.CodeConflict:
			return status.Error(codes.FailedPrecondition, linkErr.Error())
		case domain.CodeInternal:
			return status.Error(codes.Internal, linkErr.Error())
		default:
			return status.Error(codes.Internal, linkErr.Error())
		}
	}

	// Repository errors
	if errors.Is(err, repository_err.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}

	// Fallback for unknown errors
	return status.Error(codes.Internal, err.Error())
}

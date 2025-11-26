package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	domainerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/errors"
)

type InfraErrorDTO struct {
	Code      string
	Message   string
	Retryable bool
	Source    string
	Cause     error
}

func FromDomainError(source string, err error) InfraErrorDTO {
	if err == nil {
		return InfraErrorDTO{Source: source}
	}

	var derr *domainerrors.Error
	if !errors.As(err, &derr) {
		derr = domainerrors.Normalize(source, err)
	}

	return InfraErrorDTO{
		Code:      string(derr.Code()),
		Message:   derr.Error(),
		Retryable: shouldRetry(derr),
		Source:    source,
		Cause:     derr.Unwrap(),
	}
}

// ToGRPC converts the infrastructure error DTO into a gRPC status error.
func (dto InfraErrorDTO) ToGRPC() error {
	if dto.Code == "" {
		return status.Error(codes.Internal, dto.MessageOrFallback())
	}

	code := domainerrors.Code(dto.Code)

	switch code {
	case domainerrors.CodeInvalidURL:
		return status.Error(codes.InvalidArgument, dto.MessageOrFallback())
	case domainerrors.CodeMetadataNotFound:
		return status.Error(codes.NotFound, dto.MessageOrFallback())
	case domainerrors.CodeScreenshotUnavailable, domainerrors.CodeMetadataExtraction:
		if dto.Retryable {
			return status.Error(codes.Unavailable, dto.MessageOrFallback())
		}

		return status.Error(codes.FailedPrecondition, dto.MessageOrFallback())
	default:
		if dto.Retryable {
			return status.Error(codes.Unavailable, dto.MessageOrFallback())
		}

		return status.Error(codes.Internal, dto.MessageOrFallback())
	}
}

// MessageOrFallback returns the DTO message or a best-effort fallback.
func (dto InfraErrorDTO) MessageOrFallback() string {
	if dto.Message != "" {
		return dto.Message
	}

	if dto.Source != "" {
		return dto.Source + ": unknown error"
	}

	return "unknown error"
}

func shouldRetry(err *domainerrors.Error) bool {
	switch err.Code() {
	case domainerrors.CodeInvalidURL, domainerrors.CodeMetadataNotFound:
		return false
	default:
		return true
	}
}

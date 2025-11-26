package errors

import (
	"errors"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// DTO represents infrastructure-friendly view of domain link errors.
type DTO struct {
	Code      string
	Message   string
	Retryable bool
	Cause     error
	Source    string
}

// FromDomainError normalizes any error into a DTO for transport/mq adapters.
func FromDomainError(source string, err error) DTO {
	if err == nil {
		return DTO{Source: source}
	}

	var derr *domain.LinkError
	if !errors.As(err, &derr) {
		// Treat unknown errors as retryable internal failures.
		return DTO{
			Code:      string(domain.CodeInternal),
			Message:   err.Error(),
			Retryable: true,
			Cause:     err,
			Source:    source,
		}
	}

	return DTO{
		Code:      string(derr.Code()),
		Message:   derr.Error(),
		Retryable: isRetryable(derr),
		Cause:     derr.Unwrap(),
		Source:    source,
	}
}

func isRetryable(err *domain.LinkError) bool {
	switch err.Code() {
	case domain.CodeInvalidInput,
		domain.CodeNotFound,
		domain.CodePermissionDenied:
		return false
	default:
		return true
	}
}

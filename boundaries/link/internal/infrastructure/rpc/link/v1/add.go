package v1

import (
	"context"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

func (l *LinkRPC) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	// Transport-level validation — only check "empty payload"
	if in.GetLink() == nil {
		return nil, mapDomainErrorToGRPC(
			domain.NewInvalidInputError("link payload is empty"),
		)
	}

	entity, err := in.ToEntity()
	if err != nil {
		// ToEntity() → returns DomainError: InvalidInputError
		return nil, mapDomainErrorToGRPC(err)
	}

	resp, err := l.service.Add(ctx, entity)
	if err != nil {
		return nil, mapDomainErrorToGRPC(err)
	}

	return ToAddResponse(resp), nil
}

package parsers

import (
	"context"
	"errors"

	domainerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/errors"
	v1 "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
	storeerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/repository/store/error"
)

func (r *UC) Get(ctx context.Context, hash string) (*v1.Meta, error) {
	meta, err := r.MetaStore.Store.Get(ctx, hash)
	if err != nil {
		var notFound *storeerrors.MetadataNotFoundByIdError
		if errors.As(err, &notFound) {
			return nil, domainerrors.NewMetadataNotFoundError(hash, err)
		}

		return nil, domainerrors.ProcessingFailed("metadata.parser.store.get", err)
	}

	return meta, nil
}

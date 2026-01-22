package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
)

func (s *Store) MetadataUpdate(ctx context.Context, url, imageURL, description, keywords string) error {
	metadata := psql.Update("link.link_view").
		Set("image_url", imageURL).
		Set("meta_description", description).
		Set("meta_keywords", keywords).
		Where(squirrel.Eq{"url": url})

	q, args, err := metadata.ToSql()
	if err != nil {
		return err
	}

	_, err = s.client.Exec(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

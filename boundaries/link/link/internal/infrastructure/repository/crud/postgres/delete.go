package postgres

import (
	"context"
)

// Delete - delete link
func (s *Store) Delete(ctx context.Context, hash string) error {
	err := s.query.DeleteLink(ctx, hash)
	if err != nil {
		return err
	}

	return nil
}

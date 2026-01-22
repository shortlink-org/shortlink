//go:build unit || (database && postgres)

package postgres

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryGetWithMetadata(t *testing.T) {
	ctx, store, pool := newTestStore(t)

	now := time.Now().UTC().Truncate(time.Second)

	_, err := pool.Exec(ctx, `
		INSERT INTO link.link_view
			(id, url, hash, describe, image_url, meta_description, meta_keywords, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, "00000000-0000-0000-0000-000000000001", "https://example.com", "hash1", "Example",
		"https://cdn.example.com/img.png", "example page", "example keyword", now, now)
	require.NoError(t, err)

	repo, err := New(ctx, store)
	require.NoError(t, err)

	got, err := repo.Get(ctx, "hash1")
	require.NoError(t, err)
	assert.Equal(t, "hash1", got.GetHash())
	assert.Equal(t, "https://example.com", got.GetUrl().String())
	assert.Equal(t, "Example", got.GetDescribe())
	assert.Equal(t, "https://cdn.example.com/img.png", got.GetImageUrl())
	assert.Equal(t, "example page", got.GetMetaDescription())
	assert.Equal(t, "example keyword", got.GetMetaKeywords())
}

//go:build unit || (database && postgres)

package postgres

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

func TestQueryListWithMetadata(t *testing.T) {
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

	_, err = pool.Exec(ctx, `
		INSERT INTO link.link_view
			(id, url, hash, describe, image_url, meta_description, meta_keywords, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, "00000000-0000-0000-0000-000000000002", "https://example.org", "hash2", "Other",
		"https://cdn.example.org/img.png", "other page", "other keyword", now, now)
	require.NoError(t, err)

	repo, err := New(ctx, store)
	require.NoError(t, err)

	filter := &v1.FilterLink{
		URL: &v1.StringFilterInput{
			Contains: []string{"example"},
		},
	}

	list, err := repo.List(ctx, filter)
	require.NoError(t, err)
	require.Len(t, list.GetLinks(), 1)

	link := list.GetLinks()[0]
	assert.Equal(t, "hash1", link.GetHash())
	assert.Equal(t, "https://example.com", link.GetUrl().String())
	assert.Equal(t, "Example", link.GetDescribe())
	assert.Equal(t, "https://cdn.example.com/img.png", link.GetImageUrl())
	assert.Contains(t, link.GetMetaDescription(), "example")
	assert.Equal(t, "example keyword", link.GetMetaKeywords())
}

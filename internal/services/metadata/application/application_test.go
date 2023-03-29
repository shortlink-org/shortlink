//go:build unit

package metadata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	rpc "github.com/shortlink-org/shortlink/internal/services/metadata/domain/metadata/v1"
)

var metaMock = rpc.Meta{
	ImageUrl:    "",
	Keywords:    "",
	Description: "GitHub is where over 94 million developers shape the future of software, together. Contribute to the open source community, manage your Git repositories, review code like a pro, track bugs and features, power your CI/CD and DevOps workflows, and secure code before you commit it.",
}

func TestSet(t *testing.T) {
	ctx := context.Background()
	url := "https://github.com/"

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Create store
	st := &meta_store.MetaStore{}
	st.Use(ctx, log, nil)

	r, err := New(st)
	require.NoError(t, err)

	meta, err := r.Set(ctx, url)
	require.NoError(t, err, "Error get body")

	// Check content
	assert.Equal(t, metaMock.ImageUrl, meta.ImageUrl)
	assert.Equal(t, metaMock.Keywords, meta.Keywords)
	assert.Equal(t, metaMock.Description, meta.Description)
}

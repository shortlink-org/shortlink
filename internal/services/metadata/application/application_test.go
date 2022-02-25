//go:build unit

package metadata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/pkg/logger"
	rpc "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"
	"github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
)

var metaMock = rpc.Meta{
	ImageUrl:    "",
	Keywords:    "",
	Description: "GitHub is where over 73 million developers shape the future of software, together. Contribute to the open source community, manage your Git repositories, review code like a pro, track bugs and features, power your CI/CD and DevOps workflows, and secure code before you commit it.",
}

func TestSet(t *testing.T) { //nolint unused
	ctx := context.Background()
	url := "https://github.com/"

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	// Create store
	st := &meta_store.MetaStore{}
	st.Use(ctx, log, nil)

	r, err := New(st)
	assert.Nil(t, err)

	meta, err := r.Set(ctx, url)
	assert.Nil(t, err, "Error get body")

	// Check content
	assert.Equal(t, metaMock.ImageUrl, meta.ImageUrl)
	assert.Equal(t, metaMock.Keywords, meta.Keywords)
	assert.Equal(t, metaMock.Description, meta.Description)
}

package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/logger"
	rpc "github.com/batazor/shortlink/internal/metadata/domain"
	"github.com/batazor/shortlink/internal/metadata/infrastructure/store"
)

var metaMock = rpc.Meta{
	ImageURL:    "",
	Keywords:    "",
	Description: "GitHub is where over 50 million developers shape the future of software, together. Contribute to the open source community, manage your Git repositories, review code like a pro, track bugs and features, power your CI/CD and DevOps workflows, and secure code before you commit it.",
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

	r := Service{
		Store: st,
	}
	meta, err := r.Set(ctx, url)
	assert.Nil(t, err, "Error get body")

	// Check content
	assert.Equal(t, metaMock.ImageURL, meta.ImageURL)
	assert.Equal(t, metaMock.Keywords, meta.Keywords)
	assert.Equal(t, metaMock.Description, meta.Description)
}

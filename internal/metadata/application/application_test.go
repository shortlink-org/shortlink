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
	Description: "GitHub is the secure place where over 50 million developers discover, share, and build better software. Individuals and companies are streamlining DevOps and CI CD workflows, hosting and reviewing code, increasing efficiency for project management, and so much more.",
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

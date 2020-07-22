package store

import (
	"context"
	"testing"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/stretchr/testify/assert"
)

// TestLink ...
func TestLink(t *testing.T) { //nolint unused
	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	var st Store
	s := st.Use(ctx, log)

	// Init store
	assert.Nil(t, s.Init(ctx), "Error  create a new link list")

	newLink, err := link.NewURL("example.com")
	assert.Nil(t, err, "Error create a new link")

	// test add new a link
	link, err := s.Add(ctx, newLink)
	assert.Nil(t, err)

	// test get link
	link, err = s.Get(ctx, link.Hash)
	assert.Nil(t, err)
	assert.Equal(t, link.Url, newLink.Url)

	// test get links
	links, err := s.List(ctx, nil)
	assert.Nil(t, err)
	assert.Equal(t, len(links), 1)

	// test get links with using a filter
	filter := &query.Filter{Link: &query.StringFilterInput{Eq: &newLink.Url}}
	linksWithFilter, err := s.List(ctx, filter)
	assert.Nil(t, err)
	assert.Equal(t, len(linksWithFilter), 1)

	// delete link
	err = s.Delete(ctx, newLink.Hash)
	assert.Nil(t, err, "Error delete item")

	// check get after deleted
	_, err = s.Get(ctx, newLink.Hash)
	assert.NotNil(t, err, "Assert 'Not found' but get nil")
}

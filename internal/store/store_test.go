package store

import (
	"context"
	"testing"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/pkg/link"
)

// TestLink ...
func TestLink(t *testing.T) { //nolint unused
	ctx := context.Background()

	// Init logger
	conf := logger.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}
	ctx = logger.WithLogger(ctx, log)

	var st Store
	s := st.Use(ctx)

	if err = s.Init(); err != nil {
		t.Errorf("Error  create a new link list: %s", err)
	}

	newLink, err := link.NewURL("example.com")
	if err != nil {
		t.Errorf("Error create a new link: %s", err)
	}

	// test add new a link
	link, err := s.Add(newLink)
	if err != nil {
		t.Errorf("Error %s", err)
	}

	// test get link
	link, err = s.Get(link.Hash)
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if link.Url != newLink.Url {
		t.Errorf("Assert links: %s; Get %s", newLink.Url, link.Url)
	}

	// test get links
	links, err := s.List()
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if len(links) != 1 {
		t.Errorf("Assert 1 links; Get %d link(s)", len(links))
	}

	// delete link
	err = s.Delete(newLink.Hash)
	if err != nil {
		t.Errorf("Error delete item %s", err)
	}
	_, err = s.Get(newLink.Hash)
	if err == nil {
		t.Errorf("Assert 'Not found' but get nil")
	}
}

package store

import (
	"testing"

	"github.com/batazor/shortlink/pkg/link"
)

// TestLink ...
func TestLink(t *testing.T) { //nolint unused
	var st Store
	s := st.Use()

	if err := s.Init(); err != nil {
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
	if link.URL != newLink.URL {
		t.Errorf("Assert links: %s; Get %s", newLink.URL, link.URL)
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

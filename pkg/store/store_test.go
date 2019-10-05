package store

import (
	"github.com/batazor/shortlink/pkg/link"
	"testing"
)

func TestLink(t *testing.T) {
	var st Store
	s := st.Use()

	if err := s.Init(); err != nil {
		t.Errorf("Error  create a new link list: %s", err)
	}

	newLink, err := link.NewURL("example.com")
	if err != nil {
		t.Errorf("Error  create a new link: %s", err)
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

	// delete link
	err = s.Delete(newLink.Hash)
	if err != nil {
		t.Errorf("Error delete item %s", err)
	}
	link, err = s.Get(newLink.Hash)
	if err == nil {
		t.Errorf("Assert 'Not founr' but get nil")
	}
}

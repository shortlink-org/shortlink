package store

import (
	"testing"
)

func TestSQLite(t *testing.T) {
	store := SQLiteLinkList{}

	err := store.Init()
	if err != nil {
		t.Errorf("Get error: %s", err)
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(addLink)
		if err != nil {
			t.Error(err)
		}

		if link.Hash != getLink.Hash {
			t.Errorf("Assert hash - %s; Get %s hash", getLink.Hash, link.Hash)
		}
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(getLink.Hash)
		if err != nil {
			t.Error(err)
		}

		if link.Hash != getLink.Hash {
			t.Errorf("Assert hash - %s; Get %s hash", getLink.Hash, link.Hash)
		}
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List()
		if err != nil {
			t.Error(err)
		}

		if len(links) != 1 {
			t.Errorf("Assert 1 links; Get %d link(s)", len(links))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := store.Delete(getLink.Hash)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Close", func(t *testing.T) {
		err := store.Close()
		if err != nil {
			t.Error(err)
		}
	})
}

package leveldb

import (
	"testing"

	"github.com/batazor/shortlink/internal/store/mock"
)

func TestLevelDB(t *testing.T) {
	store := LevelDBLinkList{}

	err := store.Init()
	if err != nil {
		t.Errorf("Get error: %s", err)
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(mock.AddLink)
		if err != nil {
			t.Error(err)
		}

		if link.Hash != mock.GetLink.Hash {
			t.Errorf("Assert hash - %s; Get %s hash", mock.GetLink.Hash, link.Hash)
		}
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(mock.GetLink.Hash)
		if err != nil {
			t.Error(err)
		}

		if link.Hash != mock.GetLink.Hash {
			t.Errorf("Assert hash - %s; Get %s hash", mock.GetLink.Hash, link.Hash)
		}
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(nil)
		if err != nil {
			t.Error(err)
		}

		if len(links) != 1 {
			t.Errorf("Assert 1 links; Get %d link(s)", len(links))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := store.Delete(mock.GetLink.Hash)
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

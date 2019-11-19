package store

import "testing"

func TestRAM(t *testing.T) {
	store := RAMLinkList{}

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
}

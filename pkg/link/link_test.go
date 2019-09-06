package link

import "testing"

func TestLink(t *testing.T) {
	linkList, err := Init()
	if err != nil {
		t.Errorf("Error  create a new link list: %s", err)
	}

	newLink, err := NewURL("example.com")
	if err != nil {
		t.Errorf("Error  create a new link: %s", err)
	}

	// test add new a link
	err = linkList.Add(newLink)
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if len(linkList.links) != 1 {
		t.Errorf("Assert links: 1; Get %d", len(linkList.links))
	}

	// test get link
	link, err := linkList.Get(newLink)
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if link.Url != newLink.Url {
		t.Errorf("Assert links: %s; Get %s", newLink.Url, link.Url)
	}

	// delete link
	err = linkList.Delete(newLink)
	if err != nil {
		t.Errorf("Error delete item %s", err)
	}
	link, err = linkList.Get(newLink)
	if err == nil {
		t.Errorf("Error %s", err)
	}
}

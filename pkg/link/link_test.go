package link

import "testing"

func TestLink(t *testing.T) {
	linkList := LinkList{
		links: make(map[string]Link),
	}
	newLink := Link{Url: "example.com"}

	// test add new a link
	err := linkList.Add(newLink)
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
		t.Errorf("Error %s", err)
	}
	link, err = linkList.Get(newLink)
	if err == nil {
		t.Errorf("Get %s, assert error: %s", link.Url, err)
	}
}

func TestGetLink(t *testing.T) {
	linkList := LinkList{}
	link := Link{Url: "example.com"}

	linkList.Get(link)
	if len(linkList.links) != 1 {
		t.Errorf("Assert links: 1; Get %d", len(linkList.links))
	}
}

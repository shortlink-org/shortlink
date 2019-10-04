package link

import "sync"

type Link struct {
	Url      string
	Hash     string
	Describe string
}

// TODO: private fields?
type LinkList struct {
	Links map[string]Link
	Mu    sync.Mutex
}

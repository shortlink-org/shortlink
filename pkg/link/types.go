package link

import "sync"

type Link struct {
	Url      string
	Hash     string
	Describe string
}

type LinkList struct {
	links map[string]Link
	mu    sync.Mutex
}

package link

import "sync"

type Link struct {
	Url string
}

type LinkList struct {
	links map[string]Link
	mu    sync.Mutex
}

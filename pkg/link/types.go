package link

import "sync"

type Link struct {
	Url string
}

type LinkList struct {
	links []Link
	mu    sync.Mutex
}

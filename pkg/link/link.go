package link

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
)

var linkList = LinkList{
	links: make(map[string]Link),
}

func Init() (*LinkList, error) {
	return &linkList, nil
}

func (l *LinkList) Get(link Link) (*Link, error) {
	l.mu.Lock()
	response := l.links[link.Hash]
	l.mu.Unlock()

	if response.Url == "" {
		return nil, &NotFoundError{Link: link, Err: errors.New(fmt.Sprintf("Not found link: %s", link.Url))}
	}

	return &response, nil
}

func (l *LinkList) Add(link Link) (*Link, error) {
	hash := getHash([]byte(link.Url), []byte("secret"))
	link.Hash = hash[:7]

	l.mu.Lock()
	l.links[link.Hash] = link
	l.mu.Unlock()

	return &link, nil
}

func (l *LinkList) Update(link Link) (*Link, error) {
	return nil, nil
}

func (l *LinkList) Delete(link Link) error {
	l.mu.Lock()
	delete(l.links, link.Hash)
	l.mu.Unlock()

	return nil
}

func NewURL(link string) (Link, error) {
	newLink := Link{Url: link}
	return newLink, nil
}

func getHash(str, secret []byte) string {
	h := hmac.New(sha512.New, secret)
	h.Write(str)
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

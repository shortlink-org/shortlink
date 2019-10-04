package link

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

var linkList = LinkList{
	Links: make(map[string]Link),
}

func NewURL(link string) (Link, error) {
	newLink := Link{Url: link}
	return newLink, nil
}

func (l *Link) GetHash(str, secret []byte) string {
	h := hmac.New(sha512.New, secret)
	h.Write(str)
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

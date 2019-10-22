package link

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func NewURL(link string) (Link, error) {
	newLink := Link{URL: link}
	return newLink, nil
}

func (l *Link) CreateHash(str, secret []byte) string {
	h := hmac.New(sha512.New, secret)
	_, _ = h.Write(str)
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

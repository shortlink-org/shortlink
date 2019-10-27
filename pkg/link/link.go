package link

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

// NewURL return new link
func NewURL(link string) (Link, error) { // nolint unparam
	newLink := Link{URL: link}
	return newLink, nil
}

// CreateHash return hash by getting link
func (l *Link) CreateHash(str, secret []byte) string { // nolint unused
	h := hmac.New(sha512.New, secret)
	_, _ = h.Write(str) // nolint errcheck
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

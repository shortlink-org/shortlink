/*
Link entity
*/
package v1

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewURL return new link
func NewURL(link *Link) error {
	link.Hash = NewHash(link.Url)

	// Add timestamp
	link.CreatedAt = timestamppb.Now()
	link.UpdatedAt = timestamppb.Now()

	return nil
}

func NewHash(url string) string {
	return CreateHash([]byte(url), []byte("secret"))[:9]
}

// CreateHash return hash by getting link
func CreateHash(str, secret []byte) string { // nolint unused
	h := hmac.New(sha512.New, secret)
	_, _ = h.Write(str) // nolint errcheck
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

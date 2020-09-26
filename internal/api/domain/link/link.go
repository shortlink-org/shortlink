/*
Link entity
*/
package link

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"github.com/golang/protobuf/ptypes"
)

// NewURL return new link
func NewURL(link *Link) error {
	link.Hash = CreateHash([]byte(link.Url), []byte("secret"))[:9]

	// Add timestamp
	link.CreatedAt = ptypes.TimestampNow()
	link.UpdatedAt = ptypes.TimestampNow()

	return nil
}

// CreateHash return hash by getting link
func CreateHash(str, secret []byte) string { // nolint unused
	h := hmac.New(sha512.New, secret)
	_, _ = h.Write(str) // nolint errcheck
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

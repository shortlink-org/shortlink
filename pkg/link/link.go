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
func NewURL(link string) (*Link, error) { // nolint unparam
	hash := CreateHash([]byte(link), []byte("secret"))

	newLink := Link{
		Url:       link,
		Hash:      hash[:9],
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
	}

	return &newLink, nil
}

// CreateHash return hash by getting link
func CreateHash(str, secret []byte) string { // nolint unused
	h := hmac.New(sha512.New, secret)
	_, _ = h.Write(str) // nolint errcheck
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

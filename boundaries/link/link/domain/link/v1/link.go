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
	link.Hash = newHash(link.GetUrl())

	// Add timestamp
	link.CreatedAt = timestamppb.Now()
	link.UpdatedAt = timestamppb.Now()

	return nil
}

func newHash(url string) string {
	return createHash([]byte(url), []byte("secret"))[:6] //nolint:revive // ignore
}

// createHash return hash by getting link
func createHash(str, secret []byte) string {
	h := hmac.New(sha512.New, secret)
	_, _ = h.Write(str) //nolint:errcheck // ignore
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

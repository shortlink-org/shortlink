/*
Link entity
*/
package v1

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/url"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewURL return new link
func NewURL(newURL url.URL) (*link, error) {
	item := &link{
		url:  newURL,
		hash: newHash(newURL),

		// Add timestamp
		createdat: timestamppb.Now(),
		updatedat: timestamppb.Now(),
	}

	return item, nil
}

func newHash(link url.URL) string {
	return createHash([]byte(link.String()), []byte("secret"))[:15] //nolint:revive // ignore
}

// createHash return hash by getting link
func createHash(str, secret []byte) string {
	h := hmac.New(sha512.New, secret)
	_, _ = h.Write(str) //nolint:errcheck // ignore
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

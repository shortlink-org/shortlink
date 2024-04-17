//go:build gofuzz

package v1

import (
	"net/url"
)

func Fuzz(link []byte) int {
	payload, err := url.Parse(string(link))
	if err != nil {
		return -1
	}

	resp, err := NewURL(payload)
	if err != nil || len(resp.GetHash()) != 7 {
		return -1
	}

	return 1
}

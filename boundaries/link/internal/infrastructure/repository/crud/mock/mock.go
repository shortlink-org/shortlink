package mock

import (
	"fmt"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

const (
	defaultLinkURL  = "https://example.com"
	defaultDescribe = "example link"
)

// AddLink is a mock Link instance used for adding to the store.
var AddLink = mustBuildLink(defaultLinkURL, defaultDescribe)

// GetLink is a mock Link instance used for retrieving from the store.
var GetLink = mustBuildLink(defaultLinkURL, defaultDescribe)

func mustBuildLink(url, describe string) *v1.Link {
	link, err := v1.NewLinkBuilder().
		SetURL(url).
		SetDescribe(describe).
		Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build mock link: %v", err))
	}

	return link
}

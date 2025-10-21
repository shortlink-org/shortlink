package mock

import (
	"fmt"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// AddLink is a mock Link instance used for adding to the store.
var AddLink *v1.Link

// GetLink is a mock Link instance used for retrieving from the store.
var GetLink *v1.Link

func init() {
	var err error

	// Initialize AddLink using LinkBuilder
	AddLink, err = v1.NewLinkBuilder().
		SetURL("https://example.com").
		SetDescribe("example link").
		Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to build AddLink: %v", err))
	}

	// Initialize GetLink using LinkBuilder with a specific hash
	GetLink, err = v1.NewLinkBuilder().
		SetURL("https://example.com").
		SetDescribe("example link").
		Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to build GetLink: %v", err))
	}
}

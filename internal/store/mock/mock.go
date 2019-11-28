package mock

import "github.com/batazor/shortlink/pkg/link"

var (
	AddLink = link.Link{
		Url:      "https://example.com",
		Hash:     "",
		Describe: "example link",
	}

	GetLink = link.Link{
		Url:      "https://example.com",
		Hash:     "5888cab",
		Describe: "example link",
	}
)

package mock

import (
	rpc "github.com/batazor/shortlink/internal/services/metadata/domain"
)

var (
	AddMetaLink = &rpc.Meta{ // nolint unused
		Id:          "https://example.com",
		ImageURL:    "",
		Description: "example link",
		Keywords:    "",
	}

	GetMetaLink = &rpc.Meta{ // nolint unused
		Id:          "https://example.com",
		ImageURL:    "",
		Description: "example link",
		Keywords:    "",
	}
)

package mock

import (
	rpc "github.com/shortlink-org/shortlink/internal/services/metadata/domain/metadata/v1"
)

var (
	AddMetaLink = &rpc.Meta{
		Id:          "https://example.com",
		ImageUrl:    "",
		Description: "example link",
		Keywords:    "",
	}

	GetMetaLink = &rpc.Meta{
		Id:          "https://example.com",
		ImageUrl:    "",
		Description: "example link",
		Keywords:    "",
	}
)

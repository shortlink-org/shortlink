package dto

import (
	"github.com/shortlink-org/shortlink/boundaries/api/bff-web/infrastructure/http/api"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1"
)

func MakeAddLinkRequest(in api.AddLink) *v1.Link {
	var describe string

	if in.Describe != nil {
		describe = *in.Describe
	}

	return &v1.Link{
		Describe: describe,
		Url:      in.Url,
	}
}

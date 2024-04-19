package dto

import (
	"github.com/shortlink-org/shortlink/boundaries/link/bff/infrastructure/http/api"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1"
)

func MakeAddLinkRequest(in api.AddLink) *v1.Link {
	var describe string

	if in.Describe != nil {
		describe = *in.Describe
	}

	link, err := v1.NewLinkBuilder().
		SetURL(in.Url).
		SetDescribe(describe).
		Build()

	if err != nil {
		return nil
	}

	return link
}

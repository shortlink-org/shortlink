package dto

import (
	v1 "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/infrastructure/rpc/link/v1"

	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/api"
)

func MakeAddLinkRequest(in api.AddLink) *v1.Link {
	var describe string

	if in.Describe != nil {
		describe = *in.Describe
	}

	link := &v1.Link{
		Url:      in.Url,
		Describe: describe,
	}

	// Add allowed_emails if provided
	if in.AllowedEmails != nil && len(*in.AllowedEmails) > 0 {
		link.AllowedEmails = *in.AllowedEmails
	}

	return link
}

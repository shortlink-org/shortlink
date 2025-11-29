package v1

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func (l *LinkRPC) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	// Extract user email from gRPC metadata if available (from proxy service)
	// This is used for private link access control
	var userEmail string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if emails := md.Get("user-email"); len(emails) > 0 {
			userEmail = emails[0]
		}
	}

	// Pass email to usecase via context value for private link access
	// TODO: When privacy feature is implemented, usecase will check email against allowlist
	if userEmail != "" {
		ctx = context.WithValue(ctx, "user-email", userEmail)
	}

	resp, err := l.service.Get(ctx, in.GetHash())
	if err != nil {
		return nil, mapDomainErrorToGRPC(err)
	}

	return &GetResponse{
		Link: &Link{
			Url:       resp.GetUrl().String(),
			Hash:      resp.GetHash(),
			Describe:  resp.GetDescribe(),
			CreatedAt: resp.GetCreatedAt().GetTimestamp(),
			UpdatedAt: resp.GetUpdatedAt().GetTimestamp(),
		},
	}, nil
}

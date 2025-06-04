package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
)

func (l *LinkRPC) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	resp, err := l.cqrs.Get(ctx, in.GetHash())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	url := resp.GetUrl()
	return &GetResponse{
		Link: &LinkView{
			Url:             url.String(),
			Hash:            resp.GetHash(),
			Describe:        resp.GetDescribe(),
			ImageUrl:        resp.GetImageUrl(),
			MetaDescription: resp.GetMetaDescription(),
			MetaKeywords:    resp.GetMetaKeywords(),
			CreatedAt:       resp.GetCreatedAt().GetTimestamp(),
			UpdatedAt:       resp.GetUpdatedAt().GetTimestamp(),
		},
	}, nil
}

func (l *LinkRPC) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	// Parse args
	filter := v1.FilterLink{
		Url: &v1.StringFilterInput{Contains: []string{in.GetFilter()}},
	}

	resp, err := l.cqrs.List(ctx, &filter)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	links := make([]*LinkView, 0, len(resp.GetLinks()))
	for _, link := range resp.GetLinks() {
		links = append(links, &LinkView{
			Url:             link.GetUrl().String(),
			Hash:            link.GetHash(),
			Describe:        link.GetDescribe(),
			ImageUrl:        link.GetImageUrl(),
			MetaDescription: link.GetMetaDescription(),
			MetaKeywords:    link.GetMetaKeywords(),
			CreatedAt:       link.GetCreatedAt().GetTimestamp(),
			UpdatedAt:       link.GetUpdatedAt().GetTimestamp(),
		})
	}

	return &ListResponse{
		Links: &LinksView{
			Links: links,
		},
	}, nil
}

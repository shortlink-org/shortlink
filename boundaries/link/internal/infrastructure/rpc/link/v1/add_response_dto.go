package v1

import (
	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// ToAddResponse converts domain Link to AddResponse
func ToAddResponse(link *domain.Link) *AddResponse {
	return &AddResponse{
		Link: &Link{
			Url:       link.GetUrl().String(),
			Hash:      link.GetHash(),
			Describe:  link.GetDescribe(),
			CreatedAt: link.GetCreatedAt().GetTimestamp(),
			UpdatedAt: link.GetUpdatedAt().GetTimestamp(),
		},
	}
}


package v1

import (
	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// ToUpdateResponse converts domain Link to UpdateResponse
func ToUpdateResponse(link *domain.Link) *UpdateResponse {
	return &UpdateResponse{
		Link: &Link{
			Url:       link.GetUrl().String(),
			Hash:      link.GetHash(),
			Describe:  link.GetDescribe(),
			CreatedAt: link.GetCreatedAt().GetTimestamp(),
			UpdatedAt: link.GetUpdatedAt().GetTimestamp(),
		},
	}
}

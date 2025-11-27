package dto

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// Link represents the Data Transfer Object for Link JSON serialization in PostgreSQL
type Link struct {
	URL       string                `json:"url"`
	Hash      string                `json:"hash"`
	Describe  string                `json:"describe"`
	CreatedAt *timestamppb.Timestamp `json:"created_at"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at"`
}

// FromDomain converts the domain Link to the DTO Link for JSON serialization
func FromDomain(d *domain.Link) *Link {
	if d == nil {
		return nil
	}

	return &Link{
		URL:       d.GetUrl().String(),
		Hash:      d.GetHash(),
		Describe:  d.GetDescribe(),
		CreatedAt: d.GetCreatedAt().GetTimestamp(),
		UpdatedAt: d.GetUpdatedAt().GetTimestamp(),
	}
}


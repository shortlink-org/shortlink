package dto

import (
	"net/url"
	"time"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

var (
	errEmptyURL      = v1.NewInvalidInputError("mongo dto: URL cannot be empty")
	errEmptyHash     = v1.NewInvalidInputError("mongo dto: hash cannot be empty")
	errNilDomainLink = v1.NewInternalError("mongo dto: domain Link is nil")
)

// Link represents the Data Transfer Object for Link.
type Link struct {
	// URL
	Url url.URL `bson:"url" json:"url"`
	// Hash by URL + salt
	Hash string `bson:"hash" json:"hash"`
	// Describe of a link
	Describe string `bson:"describe" json:"describe"`

	// Created at
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	// Updated at
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// ToDomain converts the DTO Link to the domain Link.
func (d *Link) ToDomain() (*v1.Link, error) {
	if d.Url.String() == "" {
		return nil, errEmptyURL
	}

	if d.Hash == "" {
		return nil, errEmptyHash
	}

	// Create a new domain.Link instance
	domainLink, err := v1.NewLinkBuilder().
		SetURL(d.Url.String()).
		SetDescribe(d.Describe).
		SetCreatedAt(d.CreatedAt).
		SetUpdatedAt(d.UpdatedAt).
		Build()
	if err != nil {
		return nil, err
	}

	return domainLink, nil
}

// FromDomain converts the domain Link to the DTO Link.
func FromDomain(d *v1.Link) (*Link, error) {
	if d == nil {
		return nil, errNilDomainLink
	}

	return &Link{
		Url:       *d.GetUrl().URL,
		Hash:      d.GetHash(),
		Describe:  d.GetDescribe(),
		CreatedAt: d.GetCreatedAt().GetTime(),
		UpdatedAt: d.GetUpdatedAt().GetTime(),
	}, nil
}

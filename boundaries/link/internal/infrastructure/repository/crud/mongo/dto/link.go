package dto

import (
	"errors"
	"net/url"
	"time"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// Link represents the Data Transfer Object for Link.
type Link struct {
	// URL
	Url url.URL `json:"url" bson:"url"`
	// Hash by URL + salt
	Hash string `json:"hash" bson:"hash"`
	// Describe of a link
	Describe string `json:"describe" bson:"describe"`

	// Created at
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	// Updated at
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// ToDomain converts the DTO Link to the domain Link.
func (d *Link) ToDomain() (*v1.Link, error) {
	if d.Url.String() == "" {
		return nil, errors.New("URL cannot be empty")
	}
	if d.Hash == "" {
		return nil, errors.New("Hash cannot be empty")
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
		return nil, errors.New("domain Link is nil")
	}

	return &Link{
		Url:       *d.GetUrl().URL,
		Hash:      d.GetHash(),
		Describe:  d.GetDescribe(),
		CreatedAt: d.GetCreatedAt().GetTime(),
		UpdatedAt: d.GetUpdatedAt().GetTime(),
	}, nil
}

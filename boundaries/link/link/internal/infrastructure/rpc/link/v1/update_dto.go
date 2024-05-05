package v1

import (
	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
)

func (in *UpdateRequest) ToEntity() (*domain.Link, error) {
	link, err := domain.NewLinkBuilder().
		SetURL(in.Link.GetUrl()).
		SetDescribe(in.Link.GetDescribe()).
		Build()

	if err != nil {
		return nil, err
	}

	return link, nil
}

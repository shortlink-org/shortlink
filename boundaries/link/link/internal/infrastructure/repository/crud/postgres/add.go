package postgres

import (
	"context"
	"errors"

	"google.golang.org/protobuf/encoding/protojson"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres/schema/crud"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db/options"
)

// Add - an add link
func (s *Store) Add(ctx context.Context, source *domain.Link) (*domain.Link, error) {
	switch s.config.mode {
	case options.MODE_BATCH_WRITE:
		cb := s.config.job.Push(source)

		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, data
		case *domain.Link:
			return data, nil
		default:
			return nil, nil
		}
	case options.MODE_SINGLE_WRITE:
		data, err := s.singleWrite(ctx, source)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, nil
}

func (s *Store) singleWrite(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	dto := &v1.Link{
		Url:       in.GetUrl().String(),
		Hash:      in.GetHash(),
		Describe:  in.GetDescribe(),
		CreatedAt: in.GetCreatedAt().GetTimestamp(),
		UpdatedAt: in.GetUpdatedAt().GetTimestamp(),
	}

	// save as JSON. it doesn't make sense
	dataJson, err := protojson.Marshal(dto)
	if err != nil {
		return nil, err
	}

	links := psql.Insert("link.links").
		Columns("url", "hash", "describe", "json").
		Values(in.GetUrl().String(), in.GetHash(), in.GetDescribe(), string(dataJson))

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = s.client.Exec(ctx, q, args...)
	if err != nil {
		return nil, &v1.CreateLinkError{Link: *in}
	}

	return in, nil
}

func (s *Store) batchWrite(ctx context.Context, in *domain.Links) (*domain.Links, error) {
	links := make([]crud.CreateLinksParams, 0, len(in.GetLink()))

	// Create a new link
	list := in.GetLink()
	for key := range list {
		link := list[key].GetUrl()
		links = append(links, crud.CreateLinksParams{
			Url:      link.String(),
			Hash:     list[key].GetHash(),
			Describe: list[key].GetDescribe(),
			Json:     *list[key],
		})
	}

	_, err := s.query.CreateLinks(ctx, links)
	if err != nil {
		errs := make([]error, 0, len(list))
		for key := range list {
			errs = append(errs, &v1.CreateLinkError{Link: *list[key]})
		}

		return nil, errors.Join(errs...)
	}

	return in, nil
}

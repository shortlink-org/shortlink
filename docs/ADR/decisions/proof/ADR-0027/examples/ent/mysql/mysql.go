//go:generate go run entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema

package mysql

import (
	"context"
	"database/sql"
	"errors"

	entsql "entgo.io/ent/dialect/sql"
	uuid2 "github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shortlink-org/shortlink/docs/ADR/decisions/proof/ADR-0027/examples/ent/mysql/ent"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
)

// New store
func New(ctx context.Context, store *db.Store) (*Store, error) {
	client, ok := store.Store.GetConn().(*sql.DB)
	if !ok {
		return nil, errors.New("error get connection to MySQL")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", client)
	s := &Store{
		client: ent.NewClient(ent.Driver(drv)),
	}

	// Run the auto migration tool.
	if err := s.client.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

func (s Store) Get(ctx context.Context, id uuid2.UUID) (*domain.Link, error) {
	resp, err := s.client.Link.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Link{
		Url:       resp.URL,
		Hash:      resp.Hash,
		Describe:  resp.Describe,
		CreatedAt: timestamppb.New(resp.CreatedAt),
		UpdatedAt: timestamppb.New(resp.UpdatedAt),
	}, nil
}

func (s Store) List(ctx context.Context, filter *query.Filter) (*domain.Links, error) {
	resp, err := s.client.Link.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	links := make([]*domain.Link, 0, len(resp))
	for _, v := range resp {
		links = append(links, &domain.Link{
			Url:       v.URL,
			Hash:      v.Hash,
			Describe:  v.Describe,
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
		})
	}

	return &domain.Links{
		Link: links,
	}, nil
}

func (s Store) Add(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	err := s.client.Link.Create().
		SetURL(in.Url).
		SetHash(in.Hash).
		SetDescribe(in.Describe).
		SetJSON(*in).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s Store) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	resp, err := s.client.Link.UpdateOne(&ent.Link{
		Hash: in.Hash,
	}).
		SetURL(in.Url).
		SetHash(in.Hash).
		SetDescribe(in.Describe).
		SetJSON(*in).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.Link{
		Url:       resp.URL,
		Hash:      resp.Hash,
		Describe:  resp.Describe,
		CreatedAt: timestamppb.New(resp.CreatedAt),
		UpdatedAt: timestamppb.New(resp.UpdatedAt),
	}, nil
}

func (s Store) Delete(ctx context.Context, id string) error {
	uid, err := uuid2.FromBytes([]byte(id))
	if err != nil {
		return err
	}

	return s.client.Link.DeleteOneID(uid).Exec(ctx)
}

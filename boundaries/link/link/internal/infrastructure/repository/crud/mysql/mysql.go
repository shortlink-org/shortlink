//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate -f ./schema/sqlc.yaml

package mysql

import (
	"bytes"
	"context"
	"database/sql"
	"embed"

	"github.com/google/uuid"
	"github.com/segmentio/encoding/json"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mysql/schema/crud"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/db/mysql/migrate"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS
)

// New store
func New(ctx context.Context, store db.DB) (*Store, error) {
	client, ok := store.GetConn().(*sql.DB)
	if !ok {
		return nil, db.ErrGetConnection
	}

	s := &Store{
		client: crud.New(client),
	}

	// run migration
	err := migrate.Migration(ctx, store, migrations, "link")
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s Store) Get(ctx context.Context, hash string) (*domain.Link, error) {
	link, err := s.client.GetLinkByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	var payload domain.Link
	if json.NewDecoder(bytes.NewReader(link.Json)).Decode(&payload) != nil {
		return nil, err
	}

	return &payload, nil
}

func (s Store) List(ctx context.Context, _ *types.FilterLink) (*domain.Links, error) {
	links, err := s.client.GetLinks(ctx)
	if err != nil {
		return nil, err
	}

	resp := domain.NewLinks()
	for item := range links {
		link, err := domain.NewLinkBuilder().
			SetURL(links[item].Url).
			SetDescribe(links[item].Describe.String).
			Build()
		if err != nil {
			return nil, err
		}

		resp.Push(link)
	}

	return resp, nil
}

func (s Store) Add(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	payload, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	link := in.GetUrl()
	_, err = s.client.CreateLink(ctx, crud.CreateLinkParams{
		ID:       uuid.New(),
		Url:      link.String(),
		Hash:     in.GetHash(),
		Describe: sql.NullString{String: in.GetDescribe(), Valid: true},
		Json:     payload,
	})
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s Store) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	payload, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	link := in.GetUrl()
	_, err = s.client.UpdateLink(ctx, crud.UpdateLinkParams{
		Url:      link.String(),
		Hash:     in.GetHash(),
		Describe: sql.NullString{String: in.GetDescribe(), Valid: true},
		Json:     payload,
	})
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s Store) Delete(ctx context.Context, hash string) error {
	err := s.client.DeleteLink(ctx, hash)
	if err != nil {
		return err
	}

	return nil
}

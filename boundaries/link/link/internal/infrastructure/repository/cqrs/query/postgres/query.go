package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	v12 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link_cqrs/v1"
	v13 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func New(_ context.Context, store db.DB) (*Store, error) {
	var ok bool
	s := &Store{}

	// Set configuration
	s.client, ok = store.GetConn().(*pgxpool.Pool)
	if !ok {
		return nil, db.ErrGetConnection
	}

	return s, nil
}

// Get - get
func (s *Store) Get(ctx context.Context, id string) (*v12.LinkView, error) {
	links := psql.Select("url, hash, describe", "image_url", "meta_description", "meta_keywords").
		From("link.link_view").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}
	if rows.Err() != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	var response v12.LinkView
	for rows.Next() {
		// err = rows.Scan(&response.Url, &response.Hash, &response.Describe, &response.ImageUrl, &response.MetaDescription, &response.MetaKeywords)
		// if err != nil {
		// 	return nil, &v1.NotFoundByHashError{Hash: id}
		// }
	}

	if response.GetHash() == "" {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	return &response, nil
}

// List - list
func (s *Store) List(ctx context.Context, filter *v13.FilterLink) (*v12.LinksView, error) {
	links := psql.Select("hash, describe, ts_headline(meta_description, q, 'StartSel=<em>, StopSel=</em>') as meta_description, created_at, updated_at").
		From(fmt.Sprintf(`link.link_view, to_tsquery('%s') AS q`, filter.Url.Contains)).
		Where("make_tsvector_link_view(meta_keywords, meta_description) @@ q").
		OrderBy("ts_rank(make_tsvector_link_view(meta_keywords, meta_description), q) DESC")
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}

	response := &v12.LinksView{
		// Links: []*v12.LinkView{},
	}

	for rows.Next() {
		// var result v12.LinkView
		// var (
		// 	created_ad sql.NullTime
		// 	updated_at sql.NullTime
		// )
		// err = rows.Scan(&result.Hash, &result.Describe, &result.MetaDescription, &created_ad, &updated_at)
		// if err != nil {
		// 	return nil, &v1.NotFoundError{Link: &v1.Link{}}
		// }
		// result.CreatedAt = &timestamppb.Timestamp{Seconds: created_ad.Time.Unix(), Nanos: int32(created_ad.Time.Nanosecond())}
		// result.UpdatedAt = &timestamppb.Timestamp{Seconds: updated_at.Time.Unix(), Nanos: int32(updated_at.Time.Nanosecond())}
		//
		// response.Links = append(response.GetLinks(), &result)
	}

	return response, nil
}

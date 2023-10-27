package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	v12 "github.com/shortlink-org/shortlink/internal/services/link/domain/link_cqrs/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func New(_ context.Context, store db.DB) (*Store, error) {
	var ok bool
	s := &Store{}

	// Set configuration
	s.client, ok = store.GetConn().(*pgxpool.Pool)
	if !ok {
		return nil, fmt.Errorf("error get connection")
	}

	return s, nil
}

// Get - get
func (s *Store) Get(ctx context.Context, id string) (*v12.LinkView, error) {
	// query builder
	links := psql.Select("url, hash, describe", "image_url", "meta_description", "meta_keywords").
		From("link.link_view").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}
	if rows.Err() != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	var response v12.LinkView
	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe, &response.ImageUrl, &response.MetaDescription, &response.MetaKeywords)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
		}
	}

	if response.GetHash() == "" {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	return &response, nil
}

// List - list
func (s *Store) List(ctx context.Context, filter *query.Filter) (*v12.LinksView, error) {
	// query builder
	links := psql.Select("hash, describe, ts_headline(meta_description, q, 'StartSel=<em>, StopSel=</em>') as meta_description, created_at, updated_at").
		From(fmt.Sprintf(`link.link_view, to_tsquery('%s') AS q`, *filter.Search.Contains)).
		Where("make_tsvector_link_view(meta_keywords, meta_description) @@ q").
		OrderBy("ts_rank(make_tsvector_link_view(meta_keywords, meta_description), q) DESC").
		Limit(uint64(filter.Pagination.Limit)).
		Offset(uint64(filter.Pagination.Page * filter.Pagination.Limit))
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: query.ErrNotFound}
	}

	response := &v12.LinksView{
		Links: []*v12.LinkView{},
	}

	for rows.Next() {
		var result v12.LinkView
		var (
			created_ad sql.NullTime
			updated_at sql.NullTime
		)
		err = rows.Scan(&result.Hash, &result.Describe, &result.MetaDescription, &created_ad, &updated_at)
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: query.ErrNotFound}
		}
		result.CreatedAt = &timestamppb.Timestamp{Seconds: created_ad.Time.Unix(), Nanos: int32(created_ad.Time.Nanosecond())}
		result.UpdatedAt = &timestamppb.Timestamp{Seconds: updated_at.Time.Unix(), Nanos: int32(updated_at.Time.Nanosecond())}

		response.Links = append(response.GetLinks(), &result)
	}

	return response, nil
}

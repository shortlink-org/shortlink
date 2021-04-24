package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
	"github.com/go-redis/redis/v8"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *redis.Client
}

// Init ...
func (s *Store) Init(_ context.Context, db *db.Store) error {
	s.client = db.Store.GetConn().(*redis.Client)
	return nil
}

// Get ...
func (r *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	val, err := r.client.Get(ctx, id).Result()
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	if err = json.Unmarshal([]byte(val), &response); err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	return &response, nil
}

// List ...
func (r *Store) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
	keys := r.client.Keys(ctx, "*")
	links := []*link.Link{}

	for _, key := range keys.Val() {
		var response link.Link
		val, err := r.client.Get(ctx, key).Result()
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		if err = json.Unmarshal([]byte(val), &response); err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		links = append(links, &response)
	}

	return links, nil
}

// Add ...
func (r *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	err := link.NewURL(source)
	if err != nil {
		return nil, err
	}

	val, err := json.Marshal(source)
	if err != nil {
		return nil, &link.NotFoundError{Link: source, Err: fmt.Errorf("Failed marsharing link: %s", source.Url)}
	}

	if err = r.client.Set(ctx, source.Hash, val, 0).Err(); err != nil {
		return nil, &link.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

// Update ...
func (r *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (r *Store) Delete(ctx context.Context, id string) error {
	if err := r.client.Del(ctx, id).Err(); err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}

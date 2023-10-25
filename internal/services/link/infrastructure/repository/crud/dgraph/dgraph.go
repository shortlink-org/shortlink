package dgraph

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
)

// Link implementation of db interface
type Link struct {
	Uid      string `json:"uid,omitempty"`
	*v1.Link `json:"link,omitempty"`
	DType    []string `json:"dgraph.type,omitempty"`
}

// LinkResponse - response
type LinkResponse struct {
	Link []struct {
		*v1.Link
		Uid string `json:"uid,omitempty"`
	}
}

// Store - store struct
type Store struct {
	client *dgo.Dgraph

	log logger.Logger
}

// New store
func New(ctx context.Context, store *db.Store, log logger.Logger) (*Store, error) {
	s := &Store{
		log:    log,
		client: store.Store.GetConn().(*dgo.Dgraph),
	}

	return s, nil
}

// get - private `get` method
func (s *Store) get(ctx context.Context, id string) (*LinkResponse, error) {
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.log.ErrorWithContext(ctx, err.Error())
		}
	}()

	q := `
query all($a: string) {
	link(func: eq(hash, $a)) {
		uid
		url
		hash
		describe
		created_at
		updated_at
	}
}`

	val, err := txn.QueryWithVars(ctx, q, map[string]string{"$a": id})
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	var response LinkResponse
	if err = json.Unmarshal(val.Json, &response); err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("failed parse link: %s", id)}
	}

	return &response, nil
}

// Get public `get` method
func (s *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.log.ErrorWithContext(ctx, err.Error())
		}
	}()

	response, err := s.get(ctx, id)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	if len(response.Link) == 0 {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	return response.Link[0].Link, nil
}

// get - private `get` method
func (s *Store) list(ctx context.Context) (*LinkResponse, error) {
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.log.ErrorWithContext(ctx, err.Error())
		}
	}()

	q := `
query all {
	Link(func: has(hash)) {
		uid
		url
		hash
		describe
		created_at
		updated_at
	}
}`

	val, err := txn.QueryWithVars(ctx, q, map[string]string{})
	if err != nil {
		return nil, err
	}

	var response LinkResponse
	if errUnmarshal := json.Unmarshal(val.Json, &response); errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &response, nil
}

// List - list
func (s *Store) List(ctx context.Context, _ *query.Filter) (*v1.Links, error) {
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.log.ErrorWithContext(ctx, err.Error())
		}
	}()

	responses, err := s.list(ctx)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: query.ErrNotFound}
	}

	links := &v1.Links{
		Link: []*v1.Link{},
	}
	for _, response := range responses.Link {
		links.Link = append(links.GetLink(), &v1.Link{
			Url:      response.Url,
			Hash:     response.Hash,
			Describe: response.Describe,
		})
	}

	return links, nil
}

// Add - add
func (s *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	txn := s.client.NewTxn()
	defer func() {
		if errTxn := txn.Discard(ctx); errTxn != nil {
			s.log.ErrorWithContext(ctx, err.Error())
		}
	}()

	item := Link{
		Uid:   fmt.Sprintf(`_:%s`, source.GetHash()),
		Link:  source,
		DType: []string{"Link"},
	}

	item.Link.CreatedAt = nil
	item.Link.UpdatedAt = nil

	pb, err := protojson.Marshal(item)
	if err != nil {
		return nil, err
	}

	mu := &api.Mutation{
		SetJson:   pb,
		CommitNow: true,
		// TODO: Add condition
		// Cond: `@if(eq(len(hash), 1))`,
		// SetNquads: []byte(fmt.Sprintf(`uid(hash) <hash> "%s" .`, data.Hash)),
	}
	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("failed save link: %s", source.GetUrl())}
	}

	return source, nil
}

// Update - update
func (s *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete - delete
func (s *Store) Delete(ctx context.Context, id string) error {
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.log.ErrorWithContext(ctx, err.Error())
		}
	}()

	links, err := s.get(ctx, id)
	if err != nil {
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	if len(links.Link) == 0 {
		return nil
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	for _, delLink := range links.Link {
		dgo.DeleteEdges(mu, delLink.Uid, "hash")
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return err
	}

	return nil
}

package dgraph

import (
	"bytes"
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/segmentio/encoding/json"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

// Link implementation of db interface
type Link struct {
	Uid      string `json:"uid,omitempty"`
	*v1.Link `json:"link,omitempty"`
	DType    []string `json:"dgraph.type,omitempty"`
}

// LinkData represents the data of a link
type LinkData struct {
	*v1.Link
	Uid string `json:"uid,omitempty"`
}

// LinkResponse - response
type LinkResponse struct {
	Link []LinkData
}

// Store - store struct
type Store struct {
	client *dgo.Dgraph

	log logger.Logger
}

// New store
func New(ctx context.Context, store db.DB, log logger.Logger) (*Store, error) {
	conn, ok := store.GetConn().(*dgo.Dgraph)
	if !ok {
		return nil, db.ErrGetConnection
	}

	s := &Store{
		log:    log,
		client: conn,
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
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	var response LinkResponse
	if err = json.NewDecoder(bytes.NewReader(val.Json)).Decode(&response); err != nil {
		return nil, &v1.NotFoundByHashError{Hash: id}
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
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	if len(response.Link) == 0 {
		return nil, &v1.NotFoundByHashError{Hash: id}
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
	if errUnmarshal := json.NewDecoder(bytes.NewReader(val.Json)).Decode(&response); errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &response, nil
}

// List - list
func (s *Store) List(ctx context.Context, _ *types.FilterLink) (*v1.Links, error) {
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.log.ErrorWithContext(ctx, err.Error())
		}
	}()

	responses, err := s.list(ctx)
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}}
	}

	links := v1.NewLinks()
	for _, response := range responses.Link {
		link := response.GetUrl()
		item, err := v1.NewLinkBuilder().
			SetURL(link.String()).
			SetDescribe(response.GetDescribe()).
			Build()
		if err != nil {
			return nil, err
		}

		links.Push(item)
	}

	return links, nil
}

// Add - add
func (s *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	txn := s.client.NewTxn()
	defer func() {
		if errTxn := txn.Discard(ctx); errTxn != nil {
			s.log.ErrorWithContext(ctx, errTxn.Error())
		}
	}()

	item := Link{
		Uid:   fmt.Sprintf(`_:%s`, source.GetHash()),
		Link:  source,
		DType: []string{"Link"},
	}

	pb, err := json.Marshal(item)
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
		return nil, &v1.NotFoundError{Link: source}
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
		return &v1.NotFoundByHashError{Hash: id}
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

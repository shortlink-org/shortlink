package dgraph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
)

// DGraphLink implementation of db interface
type DGraphLink struct { // nolint unused
	Uid        string `json:"uid,omitempty"`
	*link.Link `json:"link,omitempty"`
	DType      []string `json:"dgraph.type,omitempty"`
}

// DGraphLinkResponse ...
type DGraphLinkResponse struct { // nolint unused
	Link []struct {
		*link.Link
		Uid string `json:"uid,omitempty"`
	}
}

// Store ...
type Store struct { // nolint unused
	client *dgo.Dgraph
}

// get - private `get` method
func (dg *Store) get(ctx context.Context, id string) (*DGraphLinkResponse, error) {
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
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
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response DGraphLinkResponse

	if err = json.Unmarshal(val.Json, &response); err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	return &response, nil
}

// Get public `get` method
func (dg *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
		}
	}()

	response, err := dg.get(ctx, id)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	if len(response.Link) == 0 {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return response.Link[0].Link, nil
}

// get - private `get` method
func (dg *Store) list(ctx context.Context) (*DGraphLinkResponse, error) {
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
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

	var response DGraphLinkResponse

	if err = json.Unmarshal(val.Json, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// List ...
func (dg *Store) List(ctx context.Context, _ *query.Filter) ([]*link.Link, error) {
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
		}
	}()

	responses, err := dg.list(ctx)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	var links []*link.Link
	for _, response := range responses.Link {
		links = append(links, &link.Link{
			Url:      response.Url,
			Hash:     response.Hash,
			Describe: response.Describe,
		})
	}

	return links, nil
}

// Add ...
func (dg *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	txn := dg.client.NewTxn()
	defer func() {
		if errTxn := txn.Discard(ctx); errTxn != nil {
			// TODO: use logger
			fmt.Println(errTxn.Error())
		}
	}()

	item := DGraphLink{
		Uid:   fmt.Sprintf(`_:%s`, data.Hash),
		Link:  data,
		DType: []string{"Link"},
	}

	item.Link.CreatedAt = nil
	item.Link.UpdatedAt = nil

	pb, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	mu := &api.Mutation{
		SetJson:   pb,
		CommitNow: true,
		// TODO: Add condition
		//Cond: `@if(eq(len(hash), 1))`,
		//SetNquads: []byte(fmt.Sprintf(`uid(hash) <hash> "%s" .`, data.Hash)),
	}
	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return data, nil
}

// Update ...
func (dg *Store) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (dg *Store) Delete(ctx context.Context, id string) error {
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
		}
	}()

	links, err := dg.get(ctx, id)
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
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
		return nil
	}

	return nil
}

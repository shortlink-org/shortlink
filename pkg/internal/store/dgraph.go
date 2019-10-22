package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

// DGraphLink implementation of store interface
type DGraphLink struct {
	UID string `json:"UID,omitempty"`
	link.Link
	DType []string `json:"dgraph.type,omitempty"`
}

// DGraphLinkResponse ...
type DGraphLinkResponse struct {
	Link []struct {
		link.Link
		UID string
	}
}

// DGraphLinkList ...
type DGraphLinkList struct {
	client *dgo.Dgraph
}

// Init ...
func (dg *DGraphLinkList) Init() error {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		return err
	}
	dg.client = dgo.NewDgraphClient(api.NewDgraphClient(conn))

	if err = dg.migration(); err != nil {
		return err
	}

	return nil
}

// get - private `get` method
func (dg *DGraphLinkList) get(id string) (*DGraphLinkResponse, error) {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		_ = txn.Discard(ctx)
	}()

	q := `
query all($a: string) {
	link(func: eq(hash, $a)) {
		UID
		url
		hash
		describe
	}
}`

	val, err := txn.QueryWithVars(ctx, q, map[string]string{"$a": id})
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response DGraphLinkResponse

	if err = json.Unmarshal(val.Json, &response); err != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	return &response, nil
}

// Get public `get` method
func (dg *DGraphLinkList) Get(id string) (*link.Link, error) {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		_ = txn.Discard(ctx)
	}()

	response, err := dg.get(id)
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	if len(response.Link) == 0 {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	if response.Link[0].URL == "" {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response.Link[0].Link, nil
}

// Add ...
func (dg *DGraphLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.URL), []byte("secret"))
	data.Hash = hash[:7]

	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		_ = txn.Discard(ctx)
	}()

	item := DGraphLink{
		UID:   fmt.Sprintf(`_:%s`, data.Hash),
		Link:  data,
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
		//Cond: `@if(eq(len(hash), 1))`,
		//SetNquads: []byte(fmt.Sprintf(`UID(hash) <hash> "%s" .`, data.Hash)),
	}
	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.URL)}
	}

	return &data, nil
}

// Update ...
func (dg *DGraphLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (dg *DGraphLinkList) Delete(id string) error {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		_ = txn.Discard(ctx)
	}()

	links, err := dg.get(id)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	if len(links.Link) == 0 {
		return nil
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	for _, link := range links.Link {
		dgo.DeleteEdges(mu, link.UID, "hash")
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return nil
}

// migration - init structure
func (dg *DGraphLinkList) migration() error {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		_ = txn.Discard(ctx)
	}()

	op := &api.Operation{
		Schema: `
type Link {
    url: string
    hash: string
    describe: string
}

hash: string @index(term) @lang .
`,
	}

	return dg.client.Alter(ctx, op)
}

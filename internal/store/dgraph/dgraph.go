package dgraph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

// DGraphLink implementation of store interface
type DGraphLink struct { // nolint unused
	Uid       string `json:"uid,omitempty"`
	link.Link `json:"link,omitempty"`
	DType     []string `json:"dgraph.type,omitempty"`
}

// DGraphLinkResponse ...
type DGraphLinkResponse struct { // nolint unused
	Link []struct {
		link.Link
		Uid string `json:"uid,omitempty"`
	}
}

// DGraphConfig ...
type DGraphConfig struct { // nolint unused
	URL string
}

// DGraphLinkList ...
type DGraphLinkList struct { // nolint unused
	conn   *grpc.ClientConn
	client *dgo.Dgraph
	config DGraphConfig
}

// Init ...
func (dg *DGraphLinkList) Init() error {
	var err error

	// Set configuration
	dg.setConfig()

	dg.conn, err = grpc.Dial(dg.config.URL, grpc.WithInsecure())
	if err != nil {
		return err
	}
	dg.client = dgo.NewDgraphClient(api.NewDgraphClient(dg.conn))

	if err = dg.migrate(); err != nil {
		return err
	}

	return nil
}

// Close ...
func (dg *DGraphLinkList) Close() error {
	return dg.conn.Close()
}

// Migrate - init structure
func (dg *DGraphLinkList) migrate() error { // nolint unused
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
		}
	}()

	op := &api.Operation{
		Schema: `
type Link {
    url: string
    hash: string
    describe: string
    created_at: datetime
    updated_at: datetime
}

url: string @index(term) @lang .
hash: string @index(term) @lang .
describe: string @index(term) @lang .
created_at: datetime .
updated_at: datetime .
`,
	}

	return dg.client.Alter(ctx, op)
}

// get - private `get` method
func (dg *DGraphLinkList) get(id string) (*DGraphLinkResponse, error) {
	ctx := context.Background()
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
func (dg *DGraphLinkList) Get(id string) (*link.Link, error) {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
		}
	}()

	response, err := dg.get(id)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	if len(response.Link) == 0 {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response.Link[0].Link, nil
}

// get - private `get` method
func (dg *DGraphLinkList) list() (*DGraphLinkResponse, error) {
	ctx := context.Background()
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
func (dg *DGraphLinkList) List(filter *query.Filter) ([]*link.Link, error) { // nolint unused
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
		}
	}()

	responses, err := dg.list()
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
func (dg *DGraphLinkList) Add(source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		if errTxn := txn.Discard(ctx); errTxn != nil {
			// TODO: use logger
			fmt.Println(errTxn.Error())
		}
	}()

	item := DGraphLink{
		Uid:   fmt.Sprintf(`_:%s`, data.Hash),
		Link:  *data,
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
func (dg *DGraphLinkList) Update(data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (dg *DGraphLinkList) Delete(id string) error {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			// TODO: use logger
			fmt.Println(err.Error())
		}
	}()

	links, err := dg.get(id)
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	if len(links.Link) == 0 {
		return nil
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	for _, link := range links.Link {
		dgo.DeleteEdges(mu, link.Uid, "hash")
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return nil
	}

	return nil
}

// setConfig - set configuration
func (dg *DGraphLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_DGRAPH_URI", "localhost:9080")

	dg.config = DGraphConfig{
		URL: viper.GetString("STORE_DGRAPH_URI"),
	}
}

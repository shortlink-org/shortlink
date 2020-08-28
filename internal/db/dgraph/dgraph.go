package dgraph

import (
	"context"
	"fmt"

	"github.com/spf13/viper"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/api/domain/link"
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

// Config ...
type Config struct { // nolint unused
	URL string
}

// Store ...
type Store struct { // nolint unused
	conn   *grpc.ClientConn
	client *dgo.Dgraph
	config Config
}

// Init ...
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	s.conn, err = grpc.Dial(s.config.URL, grpc.WithInsecure())
	if err != nil {
		return err
	}
	s.client = dgo.NewDgraphClient(api.NewDgraphClient(s.conn))

	if err = s.migrate(ctx); err != nil {
		return err
	}

	return nil
}

// Close ...
func (dg *Store) Close() error {
	return dg.conn.Close()
}

// Migrate - init structure
func (dg *Store) migrate(ctx context.Context) error { // nolint unused
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

// setConfig - set configuration
func (dg *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_DGRAPH_URI", "localhost:9080") // DGRAPH URI

	dg.config = Config{
		URL: viper.GetString("STORE_DGRAPH_URI"),
	}
}

package dgraph

import (
	"context"

	"github.com/spf13/viper"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/logger"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
)

// DGraphLink implementation of db interface
type DGraphLink struct { // nolint unused
	Uid      string `json:"uid,omitempty"`
	*v1.Link `json:"link,omitempty"`
	DType    []string `json:"dgraph.type,omitempty"`
}

// DGraphLinkResponse ...
type DGraphLinkResponse struct { // nolint unused
	Link []struct {
		*v1.Link
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
	logger logger.Logger
}

func New(logger logger.Logger) *Store {
	return &Store{
		logger: logger,
	}
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

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (s *Store) Close() error {
	return s.conn.Close()
}

// Migrate - init structure
func (s *Store) migrate(ctx context.Context) error { // nolint unused
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.logger.ErrorWithContext(ctx, err.Error())
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

	return s.client.Alter(ctx, op)
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_DGRAPH_URI", "localhost:9080") // DGRAPH URI

	s.config = Config{
		URL: viper.GetString("STORE_DGRAPH_URI"),
	}
}

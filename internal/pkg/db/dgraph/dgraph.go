package dgraph

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// DGraphLink implementation of db interface
type DGraphLink struct {
	Uid      string `json:"uid,omitempty"`
	*v1.Link `json:"link,omitempty"`
	DType    []string `json:"dgraph.type,omitempty"`
}

// DGraphLinkResponse ...
type DGraphLinkResponse struct { // nolint:decorder
	Link []struct {
		*v1.Link
		Uid string `json:"uid,omitempty"`
	}
}

// Config ...
type Config struct { // nolint:decorder
	URL string
}

// Store ...
type Store struct {
	logger logger.Logger
	conn   *grpc.ClientConn
	client *dgo.Dgraph
	config Config
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

	s.conn, err = grpc.Dial(s.config.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
func (s *Store) migrate(ctx context.Context) error {
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

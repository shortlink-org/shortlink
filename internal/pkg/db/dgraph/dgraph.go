package dgraph

import (
	"context"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// Link implementation of db interface
type Link struct {
	Uid      string `json:"uid,omitempty"`
	*v1.Link `json:"link,omitempty"`
	DType    []string `json:"dgraph.type,omitempty"`
}

// LinkResponse - response from DGraph
type LinkResponse struct {
	Link []struct {
		*v1.Link
		Uid string `json:"uid,omitempty"`
	}
}

// Config - config
type Config struct {
	URL string
}

// Store - store struct
type Store struct {
	log    logger.Logger
	conn   *grpc.ClientConn
	client *dgo.Dgraph
	config Config
}

func New(log logger.Logger) *Store {
	return &Store{
		log: log,
	}
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	s.conn, err = grpc.Dial(s.config.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	s.client = dgo.NewDgraphClient(api.NewDgraphClient(s.conn))

	if errMigrate := s.migrate(ctx); errMigrate != nil {
		return errMigrate
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		_ = s.close()
	}()

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// Close - close
func (s *Store) close() error {
	return s.conn.Close()
}

// Migrate - init structure
func (s *Store) migrate(ctx context.Context) error {
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.log.ErrorWithContext(ctx, err.Error())
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

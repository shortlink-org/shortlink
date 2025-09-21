package dgraph

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

// Config - config
type Config struct {
	URL string
}

// Store - store struct
type Store struct {
	log    logger.Logger
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
	// Set configuration
	s.setConfig()

	conn, err := grpc.NewClient(
		s.config.URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
	)
	if err != nil {
		return &StoreError{
			Op:      "NewClient",
			Err:     fmt.Errorf("%w: %w", ErrDgraphClient, err),
			Details: "failed to create gRPC connection",
		}
	}

	s.client = dgo.NewDgraphClient(api.NewDgraphClient(conn))

	if errMigrate := s.migrate(ctx); errMigrate != nil {
		return &StoreError{
			Op:      "migrate",
			Err:     fmt.Errorf("%w: %w", ErrDgraphMigrate, errMigrate),
			Details: "failed to migrate Dgraph schema",
		}
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		errClose := conn.Close()
		if errClose != nil {
			s.log.ErrorWithContext(ctx, errClose.Error())
		}
	}()

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// Migrate - init structure
func (s *Store) migrate(ctx context.Context) error {
	txn := s.client.NewTxn()
	defer func() {
		if err := txn.Discard(ctx); err != nil {
			s.log.ErrorWithContext(ctx, err.Error())
		}
	}()

	// TODO: use migration tool
	operation := &api.Operation{
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

	err := s.client.Alter(ctx, operation)
	if err != nil {
		return &StoreError{
			Op:      "Alter",
			Err:     err,
			Details: "failed to alter Dgraph schema",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_DGRAPH_URI", "localhost:9080") // DGRAPH URI

	s.config = Config{
		URL: viper.GetString("STORE_DGRAPH_URI"),
	}
}

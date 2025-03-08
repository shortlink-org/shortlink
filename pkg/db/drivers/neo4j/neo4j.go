package neo4j

import (
	"context"
	"fmt"
	"net/url"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/viper"
)

// Config - configuration
type Config struct {
	URI      string
	login    string
	password string
}

// Store implementation of db interface
type Store struct {
	client neo4j.DriverWithContext
	config Config
}

// Init - init connection
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	err := s.setConfig()
	if err != nil {
		return &StoreError{
			Op:      "setConfig",
			Err:     err,
			Details: "failed to set neo4j configuration",
		}
	}

	s.client, err = neo4j.NewDriverWithContext(s.config.URI, neo4j.BasicAuth(s.config.login, s.config.password, ""))
	if err != nil {
		return &StoreError{
			Op:      "init",
			Err:     ErrClientConnection,
			Details: err.Error(),
		}
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if err := s.close(ctx); err != nil {
			// We can't return the error here since we're in a goroutine,
			// but in a real application you might want to log this
			_ = err
		}
	}()

	return nil
}

// GetConn - return connection
func (s *Store) GetConn() any {
	return s.client
}

// Close - close connection
func (s *Store) close(ctx context.Context) error {
	if err := s.client.Close(ctx); err != nil {
		return &StoreError{
			Op:      "close",
			Err:     err,
			Details: "failed to close neo4j connection",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() error {
	viper.AutomaticEnv()

	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	// Neo4j 3.5, defaults to self-signed certificates, TLS on, therefore use bolt+ssc:// or neo4j+ssc://
	viper.SetDefault("STORE_NEO4J_URI", "neo4j://localhost:7687") // NEO4J URI

	uri := viper.GetString("STORE_NEO4J_URI")

	params, err := url.ParseRequestURI(uri)
	if err != nil {
		return &StoreError{
			Op:      "setConfig",
			Err:     err,
			Details: "invalid neo4j URI",
		}
	}

	password, ok := params.User.Password()
	if ok {
		s.config.password = password
	}

	s.config = Config{
		URI:   fmt.Sprintf("%s://%s", params.Scheme, params.Host),
		login: params.User.Username(),
	}

	return nil
}

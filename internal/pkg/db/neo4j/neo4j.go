package neo4j

import (
	"context"
	"fmt"
	"net/url"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/viper"
)

// Config ...
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

// Init ...
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	err := s.setConfig()
	if err != nil {
		return err
	}

	s.client, err = neo4j.NewDriverWithContext(s.config.URI, neo4j.BasicAuth(s.config.login, s.config.password, ""))
	if err != nil {
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
	return s.client.Close(context.Background())
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
		return err
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

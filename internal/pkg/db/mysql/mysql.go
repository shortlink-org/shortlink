package mysql

import (
	"context"
	"database/sql"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Init - initialize
func (s *Store) Init(_ context.Context) error {
	// Set configuration
	err := s.setConfig()
	if err != nil {
		return err
	}

	// Connect to MySQL
	if s.client, err = sql.Open("mysql", s.config.URI); err != nil {
		return err
	}

	// Check connection
	if errPing := s.client.Ping(); errPing != nil {
		return errPing
	}

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// Close - close
func (s *Store) Close() error {
	if err := s.client.Close(); err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() error {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MYSQL_URI", "shortlink:shortlink@(localhost:3306)/shortlink") // MySQL URI

	// parse uri
	uri, err := url.Parse(viper.GetString("STORE_MYSQL_URI"))
	if err != nil {
		return err
	}

	values := uri.Query()
	values.Add("parseTime", "true")

	uri.RawQuery = values.Encode()

	s.config = Config{
		URI: uri.String(),
	}

	return nil
}

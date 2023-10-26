package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Init - initialize
func (s *Store) Init(_ context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	// Connect to MySQL
	if s.client, err = sql.Open("mysql", s.config.URI); err != nil {
		return err
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
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MYSQL_URI", "shortlink:shortlink@(localhost:3306)/shortlink?parseTime=true") // MySQL URI

	s.config = Config{
		URI: viper.GetString("STORE_MYSQL_URI"),
	}
}

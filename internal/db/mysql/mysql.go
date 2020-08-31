package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

// Init ...
func (m *Store) Init(_ context.Context) error {
	var err error

	// Set configuration
	m.setConfig()

	if m.client, err = sqlx.Connect("mysql", m.config.URI); err != nil {
		return err
	}

	// Apply migration
	err = m.migrate()
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
func (m *Store) Close() error {
	return m.client.Close()
}

// Migrate ...
func (m *Store) migrate() error { // nolint unused
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS links (
			id          int NOT NULL AUTO_INCREMENT,
			url         varchar(255) NOT NULL,
			hash        varchar(255) NOT NULL,
			description text NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
	`

	if _, err := m.client.Exec(sqlStmt); err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (m *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MYSQL_URI", "shortlink:shortlink@(localhost:3306)/shortlink?parseTime=true") // MySQL URI

	m.config = Config{
		URI: viper.GetString("STORE_MYSQL_URI"),
	}
}

package mysql

import (
	"github.com/jmoiron/sqlx"
)

// MySQLConfig ...
type MySQLConfig struct { // nolint unused
	URI string
}

// MySQLLinkList implementation of store interface
type MySQLLinkList struct { // nolint unused
	client *sqlx.DB
	config MySQLConfig
}

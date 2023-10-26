package mysql

import (
	"database/sql"
)

type Store struct {
	client *sql.DB
}

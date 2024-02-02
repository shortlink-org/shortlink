package mysql

import (
	"github.com/shortlink-org/shortlink/boundaries/link/link/infrastructure/repository/crud/mysql/schema/crud"
)

type Store struct {
	client *crud.Queries
}

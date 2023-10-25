package mysql

import (
	entsql "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mysql/ent"
)

type Store struct {
	client *entsql.Client
}

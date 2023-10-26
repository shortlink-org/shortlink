package mysql

import (
	entsql "github.com/shortlink-org/shortlink/docs/ADR/decisions/proof/ADR-0027/examples/ent/mysql/ent"
)

type Store struct {
	client *entsql.Client
}

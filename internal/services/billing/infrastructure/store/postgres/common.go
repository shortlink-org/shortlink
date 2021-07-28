package postgres

import (
	"github.com/Masterminds/squirrel"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

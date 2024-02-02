package postgres

import (
	"embed"

	"github.com/Masterminds/squirrel"
)

var (
	//go:embed migrations/*.sql
	Migrations embed.FS

	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

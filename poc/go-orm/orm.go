//go:generate protoc -I./models --gotemplate_out=all=true,template_dir=template:. link.proto
package go_orm

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

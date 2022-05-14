package engine

import (
	"github.com/batazor/shortlink/pkg/shortdb/domain/query/v1"
	table "github.com/batazor/shortlink/pkg/shortdb/domain/table/v1"
	"github.com/batazor/shortlink/pkg/shortdb/engine/file"
	"github.com/batazor/shortlink/pkg/shortdb/engine/options"
)

type Engine interface {
	Exec(*v1.Query) (interface{}, error)
	Close() error

	// Table
	CreateTable(query *v1.Query) error
	DropTable(name string) error

	// Command
	Select(query *v1.Query) ([]*table.Row, error)
	Update(query *v1.Query) error
	Insert(query *v1.Query) error
	Delete(query *v1.Query) error
}

func New(name string, ops ...options.Option) (Engine, error) {
	var err error
	var engine Engine

	switch name {
	case "file":
		fallthrough
	default:
		engine, err = file.New(ops...)
		if err != nil {
			return nil, err
		}
	}

	return engine, nil
}

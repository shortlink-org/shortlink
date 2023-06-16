package engine

import (
	page "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/page/v1"
	v1 "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/query/v1"
	"github.com/shortlink-org/shortlink/internal/services/shortdb/engine/file"
	"github.com/shortlink-org/shortlink/internal/services/shortdb/engine/options"
)

type Engine interface {
	Exec(*v1.Query) (interface{}, error)
	Close() error

	// Table
	CreateTable(query *v1.Query) error
	DropTable(name string) error

	// Index
	CreateIndex(query *v1.Query) error
	DropIndex(name string) error

	// Command
	Select(query *v1.Query) ([]*page.Row, error)
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

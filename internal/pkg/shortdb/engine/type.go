package engine

import (
	"github.com/batazor/shortlink/internal/pkg/shortdb/engine/file"
	"github.com/batazor/shortlink/internal/pkg/shortdb/engine/options"
	v1 "github.com/batazor/shortlink/internal/pkg/shortdb/query/v1"
	table "github.com/batazor/shortlink/internal/pkg/shortdb/table/v1"
)

type Engine interface {
	Exec(*v1.Query) error
	Close() error

	// Table
	CreateTable(name string, fields []*table.Field) error
	DropTable(name string) error

	// Command
	Select() error
	Update() error
	Insert() error
	Delete() error
}

func New(name string, ops ...options.Option) (*Engine, error) {
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

	return &engine, nil
}

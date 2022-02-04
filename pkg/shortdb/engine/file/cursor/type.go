package cursor

import (
	"sync"

	table "github.com/batazor/shortlink/pkg/shortdb/table/v1"
)

type Cursor struct {
	wc sync.Mutex

	// table name
	Table *table.Table
	// row position
	RowId int64
	// Indicates a position one past the last element
	EndOfTable bool
}

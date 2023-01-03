package cursor

import (
	"sync"

	table "github.com/shortlink-org/shortlink/pkg/shortdb/domain/table/v1"
)

type Cursor struct {
	wc sync.RWMutex

	// table name
	Table *table.Table
	// page position
	PageId int32
	// row position
	RowId int64
	// Indicates a position one past the last element
	EndOfTable bool
}

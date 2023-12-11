package cursor

import (
	"sync"

	table "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/table/v1"
)

type Cursor struct {
	Table      *table.Table
	RowId      int64
	mu         sync.Mutex
	PageId     int32
	EndOfTable bool
}

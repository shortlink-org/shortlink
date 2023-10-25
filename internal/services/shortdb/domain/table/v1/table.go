package v1

import (
	"github.com/spf13/viper"

	index "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/index/v1"
	query "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/query/v1"
)

func New(q *query.Query) *Table {
	return &Table{
		Name:   q.GetTableName(),
		Fields: q.GetTableFields(),
		Stats: &TableStats{
			RowsCount: 0,
			PageCount: -1,
		},
		Option: &Option{
			PageSize: viper.GetInt64("SHORTDB_PAGE_SIZE"),
		},
		Index: map[string]*index.Index{},
	}
}

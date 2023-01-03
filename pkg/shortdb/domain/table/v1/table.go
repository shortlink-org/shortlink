package v1

import (
	"github.com/spf13/viper"

	index "github.com/shortlink-org/shortlink/pkg/shortdb/domain/index/v1"
	v1 "github.com/shortlink-org/shortlink/pkg/shortdb/domain/query/v1"
)

func New(query *v1.Query) *Table {
	return &Table{
		Name:   query.TableName,
		Fields: query.TableFields,
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

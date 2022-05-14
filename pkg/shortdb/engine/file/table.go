package file

import (
	"fmt"

	"github.com/spf13/viper"

	table "github.com/batazor/shortlink/pkg/shortdb/domain/table/v1"

	"github.com/batazor/shortlink/pkg/shortdb/domain/query/v1"
)

func (f *file) CreateTable(query *v1.Query) error {
	f.mc.Lock()
	defer f.mc.Unlock()

	if f.database.Tables == nil {
		f.database.Tables = make(map[string]*table.Table)
	}

	// check
	if f.database.Tables[query.TableName] != nil {
		return fmt.Errorf("at CREATE TABLE: exist table")
	}

	f.database.Tables[query.TableName] = &table.Table{
		Name:   query.TableName,
		Fields: query.TableFields,
		Stats: &table.TableStats{
			RowsCount: 0,
			PageCount: -1,
		},
		Option: &table.Option{
			PageSize: viper.GetInt64("SHORTDB_PAGE_SIZE"),
		},
	}

	return nil
}

func (f *file) DropTable(name string) error {
	// TODO implement me
	return nil
}

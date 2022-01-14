package file

import (
	"fmt"

	"github.com/spf13/viper"

	v1 "github.com/batazor/shortlink/internal/pkg/shortdb/query/v1"
	table "github.com/batazor/shortlink/internal/pkg/shortdb/table/v1"
)

func (f *file) CreateTable(query *v1.Query) error {
	f.mc.Lock()
	defer f.mc.Unlock()

	// check
	if f.database.Tables[query.TableName] != nil {
		return fmt.Errorf("at CREATE TABLE: exist table")
	}

	f.database.Tables[query.TableName] = &table.Table{
		Name:   query.TableName,
		Fields: query.TableFields,
		Stats: &table.TableStats{
			RowsCount: 0,
			PageCount: 0,
		},
		Option: &table.Option{
			PageSize: viper.GetInt64("SHORTDB_PAGE_SIZE"),
		},
	}

	return nil
}

func (f *file) DropTable(name string) error {
	//TODO implement me
	return nil
}

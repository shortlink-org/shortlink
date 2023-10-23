package file

import (
	"fmt"

	v1 "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/query/v1"
	table "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/table/v1"
)

func (f *file) CreateTable(query *v1.Query) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.database.GetTables() == nil {
		f.database.Tables = make(map[string]*table.Table)
	}

	// check
	if f.database.GetTables()[query.GetTableName()] != nil {
		return fmt.Errorf("at CREATE TABLE: exist table")
	}

	f.database.Tables[query.GetTableName()] = table.New(query)

	return nil
}

func (f *file) DropTable(name string) error {
	// TODO implement me
	return nil
}

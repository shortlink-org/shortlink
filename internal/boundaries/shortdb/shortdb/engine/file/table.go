package file

import (
	query "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/query/v1"
	table "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/table/v1"
)

func (f *File) CreateTable(q *query.Query) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.database.GetTables() == nil {
		f.database.Tables = make(map[string]*table.Table)
	}

	// check
	if f.database.GetTables()[q.GetTableName()] != nil {
		return ErrExistTable
	}

	f.database.Tables[q.GetTableName()] = table.New(q)

	return nil
}

func (f *File) DropTable(name string) error {
	// TODO implement me
	return nil
}

package file

import (
	"fmt"

	table "github.com/batazor/shortlink/internal/pkg/shortdb/table/v1"
)

func (f *file) CreateTable(name string, fields []*table.Field) error {
	f.mc.Lock()
	defer f.mc.Unlock()

	// check
	if f.database.Tables[name] != nil {
		return fmt.Errorf("at CREATE TABLE: exist table")
	}

	f.database.Tables[name] = &table.Table{
		Name:   name,
		Fields: fields,
	}

	return nil
}

func (f *file) DropTable(name string) error {
	//TODO implement me
	return nil
}

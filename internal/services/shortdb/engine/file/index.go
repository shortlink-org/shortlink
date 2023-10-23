package file

import (
	"fmt"
	"strings"

	v2 "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/index/v1"
	v1 "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/query/v1"
	"github.com/shortlink-org/shortlink/internal/services/shortdb/engine/file/index"
	parser "github.com/shortlink-org/shortlink/internal/services/shortdb/parser/v1"
)

func (f *file) CreateIndex(query *v1.Query) error {
	t := f.database.GetTables()[query.GetTableName()]

	if t.GetIndex() == nil {
		t.Index = make(map[string]*v2.Index)
	}

	// check
	for i := range query.GetIndexs() {
		if t.GetIndex()[query.GetIndexs()[i].GetName()] != nil {
			return fmt.Errorf("at CREATE INDEX: exist index %s", query.GetIndexs()[i].GetName())
		}
	}

	// create
	for i := range query.GetIndexs() {
		// create index
		t.Index[query.GetIndexs()[i].GetName()] = &v2.Index{
			Name:   query.GetIndexs()[i].GetName(),
			Type:   query.GetIndexs()[i].GetType(),
			Fields: query.GetIndexs()[i].GetFields(),
		}

		// get all values
		// TODO: use pattern iterator
		cmd, err := parser.New(fmt.Sprintf("SELECT %s from %s", strings.Join(query.GetIndexs()[i].GetFields(), ","), query.GetTableName()))
		if err != nil {
			return err
		}
		rows, err := f.Select(cmd.GetQuery())
		if err != nil { //nolint:staticcheck
			// NOTE: ignore empty table
		}

		// build index
		tree, err := index.New(t.GetIndex()[query.GetIndexs()[i].GetName()], rows)
		if err != nil {
			return err
		}

		// save to file
		payload, err := tree.Marshal()
		if err != nil {
			return err
		}

		// save date
		openFile, err := f.createFile(fmt.Sprintf("%s_%s_%s.index.json", f.database.GetName(), query.GetTableName(), query.GetIndexs()[i].GetName()))
		if err != nil {
			return err
		}

		defer func() {
			_ = openFile.Close() // #nosec
		}()

		// Write something
		err = f.writeFile(openFile.Name(), payload)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *file) DropIndex(name string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// TODO implement me
	panic("implement me")
}

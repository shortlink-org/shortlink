package file

import (
	"fmt"
	"strings"

	v2 "github.com/shortlink-org/shortlink/pkg/shortdb/domain/index/v1"
	v1 "github.com/shortlink-org/shortlink/pkg/shortdb/domain/query/v1"
	"github.com/shortlink-org/shortlink/pkg/shortdb/engine/file/index"
	parser "github.com/shortlink-org/shortlink/pkg/shortdb/parser/v1"
)

func (f *file) CreateIndex(query *v1.Query) error {
	t := f.database.Tables[query.TableName]

	if t.Index == nil {
		t.Index = make(map[string]*v2.Index)
	}

	// check
	for i := range query.Indexs {
		if t.Index[query.Indexs[i].Name] != nil {
			return fmt.Errorf("at CREATE INDEX: exist index %s", query.Indexs[i].Name)
		}
	}

	// create
	for i := range query.Indexs {
		// create index
		t.Index[query.Indexs[i].Name] = &v2.Index{
			Name:   query.Indexs[i].Name,
			Type:   query.Indexs[i].Type,
			Fields: query.Indexs[i].Fields,
		}

		// get all values
		// TODO: use pattern iterator
		cmd, err := parser.New(fmt.Sprintf("SELECT %s from %s", strings.Join(query.Indexs[i].Fields, ","), query.TableName))
		if err != nil {
			return err
		}
		rows, err := f.Select(cmd.Query)
		if err != nil { // nolint:staticcheck
			// NOTE: ignore empty table
		}

		// build index
		tree, err := index.New(t.Index[query.Indexs[i].Name], rows)
		if err != nil {
			return err
		}

		// save to file
		payload, err := tree.Marshal()
		if err != nil {
			return err
		}

		// save date
		openFile, err := f.createFile(fmt.Sprintf("%s_%s_%s.index.json", f.database.Name, query.TableName, query.Indexs[i].Name))
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

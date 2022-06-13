package file

import (
	"fmt"
	"strings"

	v2 "github.com/batazor/shortlink/pkg/shortdb/domain/index/v1"
	v1 "github.com/batazor/shortlink/pkg/shortdb/domain/query/v1"
	"github.com/batazor/shortlink/pkg/shortdb/engine/file/index"
	parser "github.com/batazor/shortlink/pkg/shortdb/parser/v1"
)

func (f *file) CreateIndex(query *v1.Query) error {
	t := f.database.Tables[query.TableName]

	if t.Index == nil {
		t.Index = make(map[string]*index.Index)
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
		rows, err := f.Select(cmd.Query)
		if err != nil {
			return err
		}

		// build index
		tree := index.New(t.Index[query.Indexs[i].Name], rows)

		// save to file
		f.database.Tables[query.TableName].Index[query.Indexs[i].Name] = tree
	}

	return nil
}

func (f *file) DropIndex(name string) error {
	f.mc.Lock()
	defer f.mc.Unlock()

	// TODO implement me
	panic("implement me")
}

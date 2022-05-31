package file

import (
	"fmt"

	index "github.com/batazor/shortlink/pkg/shortdb/domain/index/v1"
	v1 "github.com/batazor/shortlink/pkg/shortdb/domain/query/v1"
)

func (f *file) CreateIndex(query *v1.Query) error {
	f.mc.Lock()
	defer f.mc.Unlock()

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

	return nil
}

func (f *file) DropIndex(name string) error {
	f.mc.Lock()
	defer f.mc.Unlock()

	// TODO implement me
	panic("implement me")
}

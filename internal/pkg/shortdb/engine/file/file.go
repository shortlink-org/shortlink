package file

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/batazor/shortlink/internal/pkg/shortdb/engine/options"
	v1 "github.com/batazor/shortlink/internal/pkg/shortdb/query/v1"
)

type file struct {
	path string
}

func New(opts ...options.Option) (*file, error) {
	var err error
	f := &file{}

	for _, opt := range opts {
		if err := opt(f); err != nil {
			panic(err)
		}
	}

	// if not set path, set temp directory
	if f.path == "" {
		f.path, err = ioutil.TempDir(os.TempDir(), "shortdb_")
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func (f *file) Exec(query *v1.Query) error {
	switch query.Type {
	case v1.Type_TYPE_UNSPECIFIED:
		return fmt.Errorf("exec: incorret type")
	case v1.Type_TYPE_SELECT:
		return f.Select()
	case v1.Type_TYPE_UPDATE:
		return f.Update()
	case v1.Type_TYPE_INSERT:
		return f.Insert()
	case v1.Type_TYPE_DELETE:
		return f.Delete()
	case v1.Type_TYPE_CREATE_TABLE:
		return f.CreateTable(query.TableName, query.TableFields)
	case v1.Type_TYPE_DROP_TABLE:
		return f.DropTable(query.TableName)
	}

	return nil
}

func (f *file) CreateTable(name string, fields []*v1.TableField) error {
	//TODO implement me
	return nil
}

func (f *file) DropTable(name string) error {
	//TODO implement me
	return nil
}

func (f *file) Select() error {
	//TODO implement me
	return nil
}

func (f *file) Update() error {
	//TODO implement me
	return nil
}

func (f *file) Insert() error {
	//TODO implement me
	return nil
}

func (f *file) Delete() error {
	//TODO implement me
	return nil
}

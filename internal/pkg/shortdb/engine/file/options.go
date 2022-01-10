package file

import (
	"github.com/batazor/shortlink/internal/pkg/shortdb/engine/options"
)

type Option func(file *file) error

func SetPath(path string) options.Option {
	return func(o interface{}) error {
		f := o.(*file)
		f.path = path
		return nil
	}
}

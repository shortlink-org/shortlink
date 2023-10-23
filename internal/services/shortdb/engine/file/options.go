package file

import (
	"github.com/shortlink-org/shortlink/internal/services/shortdb/engine/options"
)

type Option func(file *file) error

func SetPath(path string) options.Option {
	return func(o any) error {
		f := o.(*file) //nolint:errcheck
		f.path = path

		return nil
	}
}

func SetName(name string) options.Option {
	return func(o any) error {
		f := o.(*file) //nolint:errcheck
		f.database.Name = name

		return nil
	}
}

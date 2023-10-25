package file

import (
	"github.com/shortlink-org/shortlink/internal/services/shortdb/engine/options"
)

type Option func(file *File) error

func SetPath(path string) options.Option {
	return func(o any) error {
		f := o.(*File) //nolint:errcheck // ignore
		f.path = path

		return nil
	}
}

func SetName(name string) options.Option {
	return func(o any) error {
		f := o.(*File) //nolint:errcheck // ignore
		f.database.Name = name

		return nil
	}
}

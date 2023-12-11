package file

import (
	"errors"

	"github.com/shortlink-org/shortlink/internal/boundaries/shortdb/shortdb/engine/options"
)

var ErrInvalidType = errors.New("invalid type")

type Option func(file *File) error

func SetPath(path string) options.Option {
	return func(o any) error {
		f, ok := o.(*File)
		if !ok {
			return ErrInvalidType
		}

		f.path = path

		return nil
	}
}

func SetName(name string) options.Option {
	return func(o any) error {
		f, ok := o.(*File)
		if !ok {
			return ErrInvalidType
		}

		f.database.Name = name

		return nil
	}
}

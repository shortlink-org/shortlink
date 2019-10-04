package store

import (
  "github.com/batazor/shortlink/pkg/link"
)

type Store interface {
  Get(id string) (link.Link, error)
  Add(link link.Link) (link.Link, error)
  Update(link link.Link) (link.Link, error)
  Delete(id string) error
}

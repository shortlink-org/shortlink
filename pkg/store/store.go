package store

import (
  "github.com/batazor/shortlink/pkg/api"
)

type Store interface {
  Get(id string) (api.Link, error)
  Add(link api.Link) (api.Link, error)
  Update(link api.Link) (api.Link, error)
  Delete(id string) error
}

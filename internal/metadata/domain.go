package metadata

import rpc "github.com/batazor/shortlink/internal/metadata/domain"

type Service interface {
	Get(url string) (rpc.Meta, error)
	Set(url string) (rpc.Meta, error)
}

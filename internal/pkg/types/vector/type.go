//go:build !golangci

package vector

type Type interface {
	int | int64 | uint64 | float64 | string
}

//go:build !golangci

package vector

type Type interface {
	int64 | float64 | string | int
}

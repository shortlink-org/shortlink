//go:build !golangci

package tool

type Type interface {
	int64 | float64 | string | int
}

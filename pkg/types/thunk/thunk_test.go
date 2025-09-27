package thunk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThunkFib(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "thunk")
	t.Attr("component", "types")

		t.Attr("type", "unit")
		t.Attr("package", "thunk")
		t.Attr("component", "types")
	

	cache := make([]*Thunk[int], 41) //nolint:revive // it's test

	fib := func(n int) int {
		return cache[n-1].Force() + cache[n-2].Force()
	}

	for i := range cache {
		i := i
		cache[i] = New(func() int { return fib(i) })
	}

	cache[0].o.Set(0)
	cache[1].o.Set(1)

	//nolint:revive // it's test
	assert.Equal(t, cache[40].Force(), 102334155)
}

func TestMemoizedFib(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "thunk")
	t.Attr("component", "types")

		t.Attr("type", "unit")
		t.Attr("package", "thunk")
		t.Attr("component", "types")
	
	mem := map[int]int{
		0: 0,
		1: 1,
	}

	var fib func(int) int
	fib = func(n int) int {
		if result, ok := mem[n]; ok {
			return result
		}

		result := fib(n-1) + fib(n-2)
		mem[n] = result

		return result
	}

	//nolint:revive // it's test
	assert.Equal(t, fib(40), 102334155)
}

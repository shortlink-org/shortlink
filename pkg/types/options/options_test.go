package options

import (
	"context"
	"testing"
)

func TestOption(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "options")
	t.Attr("component", "types")

	o := New[string]()

	val, err := o.Take()
	if err == nil {
		t.Fatalf("[unexpected] wanted no value out of Option[T], got: %v", val)
	}

	o.Set("hello friendos")

	_, err = o.Take()
	if err != nil {
		t.Fatalf("[unexpected] wanted no value out of Option[T], got: %v", err)
	}

	o.Clear()

	if o.IsSome() {
		t.Fatal("Option should have none, but has some")
	}
}

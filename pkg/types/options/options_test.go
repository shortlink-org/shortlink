package options

import (
	"testing"
)

func TestOption(t *testing.T) {
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

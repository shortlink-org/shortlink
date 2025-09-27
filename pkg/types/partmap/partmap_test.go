package partmap

import (
	"context"
	"fmt"
	"os"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {

	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestPartitionedMapSetAndGet(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "partmap")
	t.Attr("component", "types")

	m, err := New(&HashSumPartitioner{10}, 10)
	if err != nil {
		t.Fatalf("Failed to create PartitionedMap: %v", err)
	}

	testKey := "key1"
	testValue := "value1"

	// Test Set
	if err := m.Set(testKey, testValue); err != nil {
		t.Errorf("Failed to set value: %v", err)
	}

	// Test Get
	val, ok := m.Get(testKey)
	if !ok {
		t.Errorf("Expected key '%s' to exist", testKey)
	}
	if val != testValue {
		t.Errorf("Expected value '%s', got '%s'", testValue, val)
	}
}


func TestPartitionedMapDelete(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "partmap")
	t.Attr("component", "types")

	m, err := New(&HashSumPartitioner{10}, 10)
	if err != nil {
		t.Fatalf("Failed to create PartitionedMap: %v", err)
	}

	testKey := "key1"
	testValue := "value1"

	// Set a value
	if err := m.Set(testKey, testValue); err != nil {
		t.Errorf("Failed to set value: %v", err)
	}

	// Delete the value
	if err := m.Delete(testKey); err != nil {
		t.Errorf("Failed to delete key: %v", err)
	}

	// Try to Get the deleted value
	_, ok := m.Get(testKey)
	if ok {
		t.Errorf("Expected key '%s' to be deleted", testKey)

	}
}

func TestPartitionedMapLen(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "partmap")
	t.Attr("component", "types")

	m, err := New(&HashSumPartitioner{10}, 10)
	if err != nil {
		t.Fatalf("Failed to create PartitionedMap: %v", err)
	}

	// Initially, length should be 0
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}

	// Add some elements
	for i := range 5 {
		if err := m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i)); err != nil {
			t.Errorf("Failed to set value: %v", err)
		}
	}

	// Check length again
	if m.Len() != 5 {
		t.Errorf("Expected length 5, got %d", m.Len())
	}
}

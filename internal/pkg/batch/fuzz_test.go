//go:build fuzz
// +build fuzz

package batch

import (
	"context"
	"testing"
	"time"
)

func FuzzBatch(f *testing.F) {
	// Adding basic types
	f.Add([]byte("byte slice"))

	f.Fuzz(func(t *testing.T, input []byte) {
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		// Define a callback function
		callback := func(items []*Item) any {
			// Add assertions here to check the state of items
			return nil
		}

		// Initialize the Batch with the callback
		batch, err := New(ctx, callback)
		if err != nil {
			t.Fatalf("Failed to create batch: %v", err)
		}

		// Push the fuzzed input to the batch
		callbackChan := batch.Push(input)

		// Optionally, you can wait for the callback to complete
		select {
		case <-callbackChan:
			// Callback completed
		case <-time.After(20 * time.Second):
			// Handle timeout if the callback takes too long
			t.Fatal("Callback timed out")
		}
	})
}

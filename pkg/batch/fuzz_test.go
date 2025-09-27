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
		// Create a context with a timeout, cancel := context.WithTimeout(t.Context(), 60*time.Second)
		defer cancel()

		// Define a callback function for handling byte slices
		callback := func(items []*Item[[]byte]) error {
			// Simulate processing by sending back the item
			for _, item := range items {
				item.CallbackChannel <- item.Item
				close(item.CallbackChannel)
			}
			return nil
		}

		// Initialize the Batch with the callback
		batch, err := New(ctx, callback)
		if err != nil {
			t.Fatalf("Failed to create batch: %v", err)
		}

		// Push the fuzzed input to the batch
		callbackChan := batch.Push(input)

		// Optionally, wait for the callback to complete
		select {
		case result := <-callbackChan:
			// Assert that the input and output match
			if string(result) != string(input) {
				t.Errorf("Expected %s, got %s", input, result)
			}
		case <-time.After(20 * time.Second):
			// Handle timeout if the callback takes too long
			t.Fatal("Callback timed out")
		}
	})
}

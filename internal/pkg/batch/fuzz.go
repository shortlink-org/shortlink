//go:build fuzz
// +build fuzz

package batch

func FuzzBatch(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		// Do something with input
	})
}

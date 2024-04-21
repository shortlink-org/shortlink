package ADR_0007

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
	enc "github.com/segmentio/encoding/json"
)

var (
	payload = struct {
		Name    string    `json:"name"`
		Balance int64     `json:"ballance"`
		User    int64     `json:"user"`
		Quality int64     `json:"quality"`
		Uid     uuid.UUID `json:"uid"`
	}{
		Balance: 100,
		User:    1,
		Name:    "test",
		Uid:     uuid.New(),
		Quality: 100,
	}
)

// simple benchmark json serialization
func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(payload)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// simple benchmark segmentio/encoding serialization
func BenchmarkMarshalSegmentio(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := enc.Marshal(payload)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// simple benchmark sonic serialization
func BenchmarkMarshalSonic(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(payload)
		if err != nil {
			b.Fatal(err)
		}
	}
}

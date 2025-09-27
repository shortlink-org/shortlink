package ADR_0007

import (
	jsonv2 "encoding/json/v2"
	"testing"

	"github.com/google/uuid"
	enc "github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/require"
)

type Payload struct {
	Name    string    `json:"name"`
	Balance int64     `json:"balance"`
	User    int64     `json:"user"`
	Quality int64     `json:"quality"`
	Uid     uuid.UUID `json:"uid"`
}

var payload = Payload{
	Balance: 100,
	User:    1,
	Name:    "test",
	Uid:     mustNewV7(nil),
	Quality: 100,
}

func BenchmarkMarshalJSONv2(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, err := jsonv2.Marshal(payload); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalSegmentio(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, err := enc.Marshal(payload); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalJSONv2(b *testing.B) {
	data, _ := jsonv2.Marshal(payload)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		var out Payload
		if err := jsonv2.Unmarshal(data, &out); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalSegmentio(b *testing.B) {
	data, _ := enc.Marshal(payload)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		var out Payload
		if err := enc.Unmarshal(data, &out); err != nil {
			b.Fatal(err)
		}
	}
}

func TestUnmarshalRoundTrip(t *testing.T) {
	data, err := jsonv2.Marshal(payload)
	require.NoError(t, err)

	var got Payload
	require.NoError(t, jsonv2.Unmarshal(data, &got))
	require.Equal(t, payload, got)
}

func mustNewV7(t *testing.T) uuid.UUID {
	id, err := uuid.NewV7()
	if t != nil {
		require.NoError(t, err)
	}
	return id
}

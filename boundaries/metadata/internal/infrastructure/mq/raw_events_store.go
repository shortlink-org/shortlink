package metadata_mq

import (
	"context"
	"sync"

	"github.com/ThreeDotsLabs/watermill/message"
)

type RawEventRecord struct {
	Topic    string
	Payload  []byte
	Metadata message.Metadata
	Error    string
}

type RawEventsStore interface {
	Save(ctx context.Context, record RawEventRecord) error
}

func newInMemoryRawEventsStore() RawEventsStore {
	return &inMemoryRawEventsStore{}
}

type inMemoryRawEventsStore struct {
	mu      sync.Mutex
	records []RawEventRecord
}

func (s *inMemoryRawEventsStore) Save(_ context.Context, record RawEventRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// store copy to avoid external mutation
	copied := RawEventRecord{
		Topic:    record.Topic,
		Payload:  cloneBytes(record.Payload),
		Metadata: cloneMetadata(record.Metadata),
		Error:    record.Error,
	}

	s.records = append(s.records, copied)

	return nil
}

func cloneBytes(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}

	dst := make([]byte, len(src))
	copy(dst, src)

	return dst
}

func cloneMetadata(src message.Metadata) message.Metadata {
	if len(src) == 0 {
		return nil
	}

	dst := make(message.Metadata, len(src))
	for k, v := range src {
		dst[k] = v
	}

	return dst
}

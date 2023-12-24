package partmap

import (
	"hash/crc32"
)

// Partitioner defines the interface for partition finding strategies.
type Partitioner interface {
	Find(key string) (uint, error)
}

// HashSumPartitioner implements a Partitioner using a hash sum.
type HashSumPartitioner struct {
	partitions uint
}

func (h *HashSumPartitioner) Find(key string) (uint, error) {
	hashSum := crc32.ChecksumIEEE([]byte(key))

	return uint(hashSum) % h.partitions, nil
}

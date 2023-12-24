package partmap

import (
	"errors"
)

// PartMap is a concurrent map with separate partitions to reduce lock contention.
type PartMap struct {
	partitions []*partition
	finder     Partitioner
}

// New creates a new PartMap with the given number of partitions and Partitioner.
func New(partitioner Partitioner, partitions uint) (*PartMap, error) {
	if partitions == 0 {
		return nil, errors.New("partitions must be greater than 0")
	}

	parts := make([]*partition, partitions)
	for i := range parts {
		parts[i] = &partition{store: make(map[string]any)}
	}

	return &PartMap{
		partitions: parts,
		finder:     partitioner,
	}, nil
}

// Get retrieves a value from the map.
func (pm *PartMap) Get(key string) (any, bool) {
	partitionIndex, err := pm.finder.Find(key)
	if err != nil {
		return nil, false
	}

	return pm.partitions[partitionIndex].get(key)
}

// Set adds a key-value pair to the map.
func (pm *PartMap) Set(key string, val any) error {
	partitionIndex, err := pm.finder.Find(key)
	if err != nil {
		return err
	}

	pm.partitions[partitionIndex].set(key, val)

	return nil
}

// Delete removes a key-value pair from the map.
func (pm *PartMap) Delete(key string) error {
	partitionIndex, err := pm.finder.Find(key)
	if err != nil {
		return err
	}

	pm.partitions[partitionIndex].delete(key)

	return nil
}

// Len returns the total number of key-value pairs in the PartMap.
func (pm *PartMap) Len() int {
	total := 0
	for _, p := range pm.partitions {
		total += p.len()
	}
	return total
}

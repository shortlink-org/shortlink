package v1

// Links aggregates link entries.
type Links struct {
	items []*Link
}

// NewLinks creates an empty Links collection.
func NewLinks() *Links {
	return &Links{
		items: make([]*Link, 0),
	}
}

// GetLinks returns a copy of the stored links slice to preserve invariants.
func (m *Links) GetLinks() []*Link {
	if m == nil || len(m.items) == 0 {
		return []*Link{}
	}

	copied := make([]*Link, len(m.items))
	copy(copied, m.items)

	return copied
}

// Count returns the number of stored links.
func (m *Links) Count() int {
	if m == nil {
		return 0
	}

	return len(m.items)
}

// Push adds new Link entries to the collection, filtering nil values.
func (l *Links) Push(link ...*Link) {
	if l == nil {
		return
	}

	for _, item := range link {
		if item == nil {
			continue
		}
		l.items = append(l.items, item)
	}
}

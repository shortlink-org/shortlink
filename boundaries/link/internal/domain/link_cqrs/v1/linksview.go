package v1

// LinksView aggregates CQRS link projections.
type LinksView struct {
	links []*LinkView
}

// GetLinks returns the value of the links field.
func (m *LinksView) GetLinks() []*LinkView {
	return m.links
}

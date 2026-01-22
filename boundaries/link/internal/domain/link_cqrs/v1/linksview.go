package v1

// LinksView aggregates CQRS link projections.
type LinksView struct {
	links []*LinkView
}

// NewLinksView returns an empty LinksView.
func NewLinksView() *LinksView {
	return &LinksView{links: []*LinkView{}}
}

// AddLink appends a link view to the list.
func (m *LinksView) AddLink(link *LinkView) {
	m.links = append(m.links, link)
}

// GetLinks returns the value of the links field.
func (m *LinksView) GetLinks() []*LinkView {
	return m.links
}

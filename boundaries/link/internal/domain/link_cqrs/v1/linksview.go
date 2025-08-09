package v1

// Links
type LinksView struct {
	// Links
	links []*LinkView
}

// GetLinks returns the value of the links field.
func (m *LinksView) GetLinks() []*LinkView {
	return m.links
}

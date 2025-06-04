package v1

// Sitemap is a structure of <sitemap>
type Sitemap struct {
	// Url is a structure of <url> in <sitemap>
	url []*Url `xml:"url"`
}

// GetUrl returns the value of the url field.
func (m *Sitemap) GetUrl() []*Url {
	return m.url
}

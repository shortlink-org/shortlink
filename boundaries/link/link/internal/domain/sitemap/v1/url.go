package v1

// URL is a structure of <url> in <sitemap>
type Url struct {
	// loc is a structure of <loc> in <url>
	loc string `xml:"loc"`
	// lastmod is a structure of <lastmod> in <url>
	lastmod string `xml:"lastmod"`
	// changefreq is a structure of <changefreq> in <url>
	changefreq string `xml:"changefreq"`
	// priority is a structure of <priority> in <url>
	priority float32 `xml:"priority"`
}

// GetLoc returns the value of the loc field.
func (m *Url) GetLoc() string {
	return m.loc
}

// GetLastMod returns the value of the lastmod field.
func (m *Url) GetLastMod() string {
	return m.lastmod
}

// GetChangeFreq returns the value of the changefreq field.
func (m *Url) GetChangeFreq() string {
	return m.changefreq
}

// GetPriority returns the value of the priority field.
func (m *Url) GetPriority() float32 {
	return m.priority
}

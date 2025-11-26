package v1

// Url is a structure of <url> in <sitemap>
type Url struct {
	// Loc is a structure of <loc> in <url>
	Loc string `xml:"loc"`
	// LastMod is a structure of <lastmod> in <url`
	LastMod string `xml:"lastmod"`
	// ChangeFreq is a structure of <changefreq> in <url>
	ChangeFreq string `xml:"changefreq"`
	// Priority is a structure of <priority> in <url>
	Priority float32 `xml:"priority"`
}

// GetLoc returns the value of the loc field.
func (m *Url) GetLoc() string {
	return m.Loc
}

// GetLastMod returns the value of the lastmod field.
func (m *Url) GetLastMod() string {
	return m.LastMod
}

// GetChangeFreq returns the value of the changefreq field.
func (m *Url) GetChangeFreq() string {
	return m.ChangeFreq
}

// GetPriority returns the value of the priority field.
func (m *Url) GetPriority() float32 {
	return m.Priority
}

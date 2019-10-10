package link

type Link struct {
	Url      string `json:"url,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Describe string `json:"describe,omitempty"`
}

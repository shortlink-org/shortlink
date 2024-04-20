package v1

import (
	"net/url"
)

type Url struct {
	*url.URL
}

func NewUrl(in string) (Url, error) {
	resp, err := url.Parse(in)
	if err != nil {
		return Url{}, err
	}

	return Url{resp}, nil
}

func (m *Url) GetUrl() *url.URL {
	return m.URL
}

func (m *Url) String() string {
	return m.URL.String()
}

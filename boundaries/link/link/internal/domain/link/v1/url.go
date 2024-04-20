package v1

import (
	"encoding/json"
	"net/url"
)

type Url struct {
	*url.URL
}

func NewUrl(in string) (Url, error) {
	resp, err := url.ParseRequestURI(in)
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

// MarshalJSON implements the json.Marshaler interface.
func (m Url) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *Url) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	m.URL = u
	return nil
}

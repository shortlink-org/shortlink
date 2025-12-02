package url

import (
	"net/url"

	"github.com/segmentio/encoding/json"
)

// URL is a Value Object representing a URL.
type URL struct {
	*url.URL
}

// NewURL creates a new URL Value Object from a string.
func NewURL(in string) (URL, error) {
	resp, err := url.ParseRequestURI(in)
	if err != nil {
		return URL{}, err
	}

	return URL{resp}, nil
}

// GetURL returns the underlying *url.URL.
func (u *URL) GetURL() *url.URL {
	return u.URL
}

// String returns the string representation of the URL.
func (u *URL) String() string {
	return u.URL.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (u URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *URL) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := url.Parse(s)
	if err != nil {
		return err
	}

	u.URL = parsed

	return nil
}



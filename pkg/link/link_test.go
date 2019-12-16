package link

import "testing"

func TestNewURL(t *testing.T) {
	URL := "http://test.com"
	newUrl, err := NewURL("http://test.com")
	if err != nil {
		t.Errorf("Assert Success. Got: %s", err)
	}

	if newUrl.Url != URL {
		t.Errorf("Assert %s. Got: %s", URL, newUrl.Url)
	}
}

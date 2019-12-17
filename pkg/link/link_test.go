package link

import "testing"

func TestNewURL(t *testing.T) {
	URL := "http://test.com"

	t.Run("create new", func(t *testing.T) {
		newUrl, err := NewURL(URL)

		if err != nil {
			t.Errorf("Assert nil. Got: %s", err)
		}

		if newUrl.Url != URL {
			t.Errorf("Assert %s. Got: %s", URL, newUrl.Url)
		}
	})

	t.Run("create hash", func(t *testing.T) {
		newUrl, err := NewURL(URL)

		if err != nil {
			t.Errorf("Assert nil. Got: %s", err)
		}

		success := "99699cbfa9614160a94114f527f5501fd97edeaa767db24bb8581789e18a9c1f0f671ee525a9404abc4a8d015315f773dd214175a9c50ac6cda1d934f75fc1e8"
		response := newUrl.CreateHash([]byte("hello world"), []byte("solt"))
		if success != response {
			t.Errorf("Assert %s. Got: %s", success, response)
		}
	})
}

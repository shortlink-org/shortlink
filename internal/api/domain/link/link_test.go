package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestNewURL(t *testing.T) {
	URL := "http://test.com"

	t.Run("create new", func(t *testing.T) {
		newUrl, err := NewURL(URL)

		assert.Nil(t, err, "Assert nil. Got: %s", err)
		assert.Equal(t, newUrl.Url, URL)
	})

	t.Run("create hash", func(t *testing.T) {
		success := "99699cbfa9614160a94114f527f5501fd97edeaa767db24bb8581789e18a9c1f0f671ee525a9404abc4a8d015315f773dd214175a9c50ac6cda1d934f75fc1e8"
		response := CreateHash([]byte("hello world"), []byte("solt"))
		assert.Equal(t, success, response)
	})
}

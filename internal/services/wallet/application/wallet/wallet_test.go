package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWalletClient(t *testing.T) {
	_, err := NewWalletClient()
	assert.NoError(t, err)
}

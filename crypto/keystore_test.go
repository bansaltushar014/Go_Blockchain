package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
	W, err := CreateWallet()
	assert.NoError(t, err)
	assert.NotEmpty(t, W)
}

func TestSignAndVerify(t *testing.T) {
	W, err := CreateWallet()
	assert.NoError(t, err)
	assert.NotEmpty(t, W)
	h, r, s := W.SignMsg([]byte("test this!"))
	assert.NotEmpty(t, h)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, s)
	b := Verify(h, r, s, W.PublicKey)
	assert.True(t, b)
}

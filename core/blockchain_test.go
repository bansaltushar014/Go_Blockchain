package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	b := CreateGenesisBlock()
	bl, err := NewBlockchain(b)
	assert.Nil(t, err)
	assert.NotEmpty(t, bl)
}

func TestAddBlocktoBlockchain(t *testing.T) {
	b := CreateGenesisBlock()
	bl, err := NewBlockchain(b)
	assert.Nil(t, err)
	x := 3
	var prevHash [32]uint8 = b.Hash()
	for i := 1; i <= x; i++ {
		bk := randomBlockCreationWithSignature(i, prevHash)
		prevHash = bk.Hash()
		err = bl.AddBlock(bk)
		assert.Nil(t, err)
		assert.Equal(t, bk.Header, bl.GetHeader())
	}
	assert.Equal(t, len(bl.header), x+1)
}

func TestValidateBlock(t *testing.T) {
	b := CreateGenesisBlock()
	bl, err := NewBlockchain(b)
	assert.Nil(t, err)
	var prevHash [32]uint8 = b.Hash()
	bk := randomBlockCreation(1, prevHash)
	err = bl.AddBlock(bk)
	assert.Nil(t, err)
}

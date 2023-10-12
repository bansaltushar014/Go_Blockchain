package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenesisBlock(t *testing.T) {
	// assert.Empty(t, B)
	B := CreateGenesisBlock()
	assert.NotEmpty(t, B)
}

func TestAddTransactions(t *testing.T) {
	B := CreateGenesisBlock()
	ok := B.addTransactions([]byte("Second Tx"))
	assert.NoError(t, ok)
	ok2 := B.addTransactions([]byte("Third Tx"))
	assert.NoError(t, ok2)
	ok3 := B.addTransactions([]byte("Forth Tx"))
	assert.NoError(t, ok3)
	ok4 := B.addTransactions([]byte("Fifth Tx"))
	assert.NoError(t, ok4)

	// Checking error on 6th Transaction
	ok5 := B.addTransactions([]byte("Sixth Tx"))
	assert.Error(t, ok5)
}

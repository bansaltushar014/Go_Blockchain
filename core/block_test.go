package core

import (
	"testing"
	"time"

	cryptoX "github.com/bansaltushar014/go-blockchain-l2/crypto"
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

func TestBlockSignndVerify(t *testing.T) {
	B := CreateGenesisBlock()
	ok := B.addTransactions([]byte("Second Tx"))
	assert.NoError(t, ok)
	W, _ := cryptoX.CreateWallet()
	B.SignBlock(W.PrivateKey)
	assert.NotNil(t, B.Signature)
	res := B.Verify()
	assert.True(t, res)
}

func createRandomTransaction() *Transaction {
	tx := excodeTx([]byte(time.Now().String()))
	transaction := Transaction{
		data: tx,
	}
	return &transaction
}

func createRandomTransactionWithSignature() *Transaction {
	W, _ := cryptoX.CreateWallet()
	tx := excodeTx([]byte(time.Now().String()))
	transaction := Transaction{
		data: tx,
	}
	transaction.signTx(W)
	return &transaction
}

func randomBlockCreation(height int, prevHash [32]uint8) *Block {
	tx := excodeTx([]byte(time.Now().String()))
	transaction := Transaction{
		data: tx,
	}
	block := Block{
		Header: &Header{
			version:       []byte("v1"),
			DataHash:      [32]uint8{},
			TimeStamp:     time.Now(),
			Height:        height,
			PrevBlockHash: prevHash,
		},
		Transactions: []Transaction{transaction},
	}
	return &block
}

func randomBlockCreationWithSignature(height int, prevHash [32]uint8) *Block {
	transaction := *createRandomTransactionWithSignature()
	transaction2 := *createRandomTransactionWithSignature()
	transaction3 := *createRandomTransactionWithSignature()

	block := Block{
		Header: &Header{
			version:       []byte("v1"),
			DataHash:      [32]uint8{},
			TimeStamp:     time.Now(),
			Height:        height,
			PrevBlockHash: prevHash,
		},
		Transactions: []Transaction{transaction, transaction2, transaction3},
	}
	W, _ := cryptoX.CreateWallet()
	block.SignBlock(W.PrivateKey)
	return &block
}

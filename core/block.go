package core

import (
	"errors"
	"time"
)

type Header struct {
	version   []byte
	prevBlock []byte
	TimeStamp time.Time
	Height    int
	Nonce     []byte
}

type Block struct {
	*Header
	Transactions []Transaction
}

const txLimit = 5 // Max 5 tx can go inside a block

func CreateGenesisBlock() *Block {
	tx := excodeTx([]byte("First Transaction"))
	transaction := Transaction{
		data: tx,
	}
	block := Block{
		Header: &Header{
			version:   []byte("v1"),
			prevBlock: []byte(""),
			TimeStamp: time.Now(),
			Height:    0,
			Nonce:     []byte(""),
		},
		Transactions: []Transaction{transaction},
	}
	return &block
}

func (b *Block) addTransactions(data []byte) error {
	if len(b.Transactions) >= 5 {
		return errors.New("Cant push more transaction")
	}

	tx := excodeTx(data)
	transaction := Transaction{
		data: tx,
	}
	b.Transactions = append(b.Transactions, transaction)
	return nil
}

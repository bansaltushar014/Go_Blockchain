package core

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"time"

	cryptoX "github.com/bansaltushar014/go-blockchain-l2/crypto"
)

type Hasht [32]uint8

type Header struct {
	version       []byte
	DataHash      [32]uint8
	TimeStamp     time.Time
	Height        int
	PrevBlockHash [32]uint8
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    cryptoX.PublicKey
	Signature    *cryptoX.Signature
}

const txLimit = 5 // Max 5 tx can go inside a block

func CreateGenesisBlock() *Block {
	tx := excodeTx([]byte("First Transaction"))
	transaction := Transaction{
		data: tx,
	}
	block := Block{
		Header: &Header{
			version:       []byte("v1"),
			DataHash:      [32]uint8{},
			TimeStamp:     time.Now(),
			Height:        0,
			PrevBlockHash: [32]uint8{},
		},
		Transactions: []Transaction{transaction},
	}
	return &block
}

func CreateRandomBlock() *Block {
	tx := excodeTx([]byte(time.Now().String()))
	transaction := Transaction{
		data: tx,
	}
	block := Block{
		Header: &Header{
			version:       []byte("v1"),
			DataHash:      [32]uint8{},
			TimeStamp:     time.Now(),
			Height:        0,
			PrevBlockHash: [32]uint8{},
		},
		Transactions: []Transaction{transaction},
	}
	return &block
}

func (b *Block) addTransactions(data []byte) error {
	if len(b.Transactions) >= 5 {
		return errors.New("Cant push more transaction")
	}

	t := Trans{}
	tx := t.CreateRandomTx(data)
	b.Transactions = append(b.Transactions, *tx)
	return nil
}

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)

	return buf.Bytes()
}

func (b *Block) SignBlock(privateKey *ecdsa.PrivateKey) error {
	blockBytes := b.Header.Bytes()
	w := cryptoX.Wallet{PrivateKey: privateKey}
	_, r, s := w.SignMsg(blockBytes)
	b.Signature = &cryptoX.Signature{r, s}
	b.Validator = cryptoX.PublicKey{&privateKey.PublicKey}
	return nil
}

func (b *Block) Verify() bool {
	if b.Signature == nil {
		return false
	}

	for _, v := range b.Transactions {
		ok := v.verify()
		if ok != true {
			return ok
		}
	}
	return true
}

func (b *Block) Hash() Hasht {
	header := b.Header
	h := sha256.Sum256(header.Bytes())
	return Hasht(h)
}

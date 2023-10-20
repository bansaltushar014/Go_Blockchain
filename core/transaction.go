package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"

	cryptoX "github.com/bansaltushar014/go-blockchain-l2/crypto"
)

type Transaction struct {
	Data      []byte
	publicKey cryptoX.PublicKey
	signature *cryptoX.Signature
	hash      [32]uint8
	firstSeen time.Time
}

// This is hash function , not encode function. Change it in l2 and l3
func ExcodeTx(data []byte) []byte {
	hasher := sha256.New()
	_, err := hasher.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return []byte("Failed to Hash!")
	}
	hashSum := hasher.Sum(nil)
	return hashSum
}

func (tx *Transaction) Hash(data []byte) [32]uint8 {
	h := sha256.Sum256(Bytes(data))
	tx.hash = [32]uint8(h)
	return tx.hash
}

func Bytes(data []byte) []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(data)

	return buf.Bytes()
}

func NewTransaction() *Transaction {
	return createTransaction([]byte("First Transaction!"))
}

type Trans struct{}

func (t *Trans) CreateRandomTx(data []byte) *Transaction {
	return createTransaction(data)
}

func createTransaction(data []byte) *Transaction {
	w, _ := cryptoX.CreateWallet()
	_, r, s := w.SignMsg([]byte(data))
	return &Transaction{
		Data:      data,
		publicKey: cryptoX.PublicKey{w.PublicKey},
		signature: &cryptoX.Signature{r, s},
	}
}

func (tx *Transaction) signTx(w *cryptoX.Wallet) error {
	// hash := sha256.Sum256(tx.data)
	_, r, s := w.SignMsg(tx.Data)
	tx.signature = &cryptoX.Signature{r, s}
	tx.publicKey = cryptoX.PublicKey{w.PublicKey}
	return nil
}

func (tx *Transaction) verify() bool {
	ok := cryptoX.Verify(tx.Data, tx.signature.R, tx.signature.S, tx.publicKey.Key)
	return ok
}

func (tx *Transaction) SetFirstSeen(t time.Time) {
	tx.firstSeen = t
}

func (tx *Transaction) FirstSeen() time.Time {
	return tx.firstSeen
}

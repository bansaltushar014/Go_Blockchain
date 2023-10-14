package core

import (
	"crypto/sha256"
	"fmt"

	cryptoX "github.com/bansaltushar014/go-blockchain-l2/crypto"
)

type Transaction struct {
	data      []byte
	publicKey cryptoX.PublicKey
	signature *cryptoX.Signature
}

// This is hash function , not encode function. Change it in l2
func excodeTx(data []byte) []byte {
	hasher := sha256.New()
	_, err := hasher.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return []byte("Failed to Hash!")
	}
	hashSum := hasher.Sum(nil)
	return hashSum
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
		data:      data,
		publicKey: cryptoX.PublicKey{w.PublicKey},
		signature: &cryptoX.Signature{r, s},
	}
}

func (tx *Transaction) signTx(w *cryptoX.Wallet) error {
	// hash := sha256.Sum256(tx.data)
	_, r, s := w.SignMsg(tx.data)
	tx.signature = &cryptoX.Signature{r, s}
	tx.publicKey = cryptoX.PublicKey{w.PublicKey}
	return nil
}

func (tx *Transaction) verify() bool {
	ok := cryptoX.Verify(tx.data, tx.signature.R, tx.signature.S, tx.publicKey.Key)
	return ok
}

package core

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math/big"

	cryptoX "github.com/bansaltushar014/go-blockchain-l2/crypto"
)

type PublicKey struct {
	key *ecdsa.PublicKey
}

type Signature struct {
	r, s *big.Int
}

type Transaction struct {
	data      []byte
	publicKey PublicKey
	signature *Signature
}

func excodeTx(data []byte) []byte {
	fmt.Println("fmt")
	hasher := sha256.New()
	_, err := hasher.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return []byte("Failed to Hash!")
	}
	hashSum := hasher.Sum(nil)
	return hashSum
}

// Start from here where you are stuck, you need to move that package from this package to other side
func (tx *Transaction) signTx(w *cryptoX.Wallet) *Transaction {
	// hash := sha256.Sum256(tx.data)
	_, r, s := w.SignMsg(string(tx.data))
	tx.signature = &Signature{r, s}
	tx.publicKey = PublicKey{w.PublicKey}
	return tx
}

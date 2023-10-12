package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    []byte
}

func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	// Generate a new private key using the P-256 elliptic curve
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Get the corresponding public key from the private key
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

func generateEthereumAddress(publicKeyBytes []byte) (string, error) {
	// Step 2: Perform Keccak-256 hashing on the public key
	hash := crypto.Keccak256(publicKeyBytes)

	// Step 3: Take the last 20 bytes of the hash to get the Ethereum address
	address := hash[len(hash)-20:]

	// Convert the address to hexadecimal format
	addressHex := fmt.Sprintf("%x", address)

	// Add "0x" prefix to the address
	addressWithPrefix := "0x" + addressHex

	return addressWithPrefix, nil
}

func CreateWallet() (*Wallet, error) {
	privKey, pubKey, err := generateKeyPair()
	if err != nil {
		return nil, errors.New("Failed to create Key Pair")
	}

	// publicKeyBytes := crypto.FromHex(pubKey)
	data := elliptic.Marshal(pubKey, pubKey.X, pubKey.Y)

	address, err := generateEthereumAddress(data)
	if err != nil {
		// fmt.Println("Error:", err)
		return nil, errors.New(err.Error())
	}

	return &Wallet{
		PrivateKey: privKey,
		PublicKey:  pubKey,
		Address:    []byte(address),
	}, nil
}

func (w *Wallet) SignMsg(msg string) ([32]byte, *big.Int, *big.Int) {
	message := []byte(msg)
	hash := sha256.Sum256(message)
	r, s, err := ecdsa.Sign(rand.Reader, w.PrivateKey, hash[:])
	if err != nil {
		fmt.Println("Error signing the message:", err)
	}
	return hash, r, s
}

func Verify(hash [32]byte, r *big.Int, s *big.Int, pubkey *ecdsa.PublicKey) bool {
	valid := ecdsa.Verify(pubkey, hash[:], r, s)
	if valid {
		return true
	} else {
		return false
	}
}
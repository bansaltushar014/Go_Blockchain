package core

import (
	"errors"
	"sync"

	"crypto/sha256"

	"github.com/sirupsen/logrus"
)

type Blockchain struct {
	store     Storage
	header    []*Header
	validator Validator
	lock      sync.RWMutex
}

// Storage
type Storage interface {
	Put(*Block) error
}

type MemoryStore struct {
}

func NewMemorystore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Put(b *Block) error {
	return nil
}

// Validator
type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (blv *BlockValidator) ValidateBlock(b *Block) error {
	if !blv.bc.HasBlock(b.Height) {
		return errors.New("Validation Failed!")

	}

	// Validate if prev block hash and current hash is same or not!
	h := blv.bc.GetHeader()
	hash := sha256.Sum256(h.Bytes())
	if hash != b.PrevBlockHash {
		return errors.New("Prev Hash mismateched!")
	}

	// Mismated
	if len(blv.bc.header) != b.Header.Height {
		return errors.New("Height Mismated")
	}

	// verify all the transaction exists inside the block
	if !b.Verify() {
		return errors.New("Block Verification Failed!")
	}

	return nil
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		header: []*Header{},
		store:  &MemoryStore{},
	}
	// bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)

	return bc, err
}

func (bl *Blockchain) GetHeader() *Header {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	return bl.header[bl.Height()]
}

func (bl *Blockchain) SetValidator(Val Validator) {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	bl.validator = Val
}

func (bl *Blockchain) AddBlock(b *Block) error {
	// Below strategy failing in while testing.
	// if err := bl.validator.ValidateBlock(b); err != nil {
	// 	return err
	// }
	bl.lock.Lock()
	defer bl.lock.Unlock()
	//check block already exists
	blv := BlockValidator{bl}
	if err := blv.ValidateBlock(b); err != nil {
		return err
	}
	bl.header = append(bl.header, b.Header)
	return nil
}

func (bl *Blockchain) HasBlock(height int) bool {
	return height <= len(bl.header)
}

func (bl *Blockchain) Height() int {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	return len(bl.header) - 1
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.lock.Lock()
	defer bc.lock.Unlock()
	bc.header = append(bc.header, b.Header)

	logrus.WithFields(logrus.Fields{
		"height": b.Height,
		"hash":   b.Hash(),
	}).Info("adding new block")

	return bc.store.Put(b)
}

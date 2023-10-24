package network

import (
	"sort"

	"github.com/bansaltushar014/go-blockchain-l2/core"
)

type TxMapSorter struct {
	transactions []*core.Transaction
}

type TxPool struct {
	transactions map[[32]uint8]*core.Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[[32]uint8]*core.Transaction),
	}
}

func (txp *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(tx.Data)
	txp.transactions[hash] = tx
	return nil
}

func (txp *TxPool) Has(hash [32]uint8) bool {
	_, ok := txp.transactions[hash]
	return ok
}

func (txp *TxPool) Len() int {
	return len(txp.transactions)
}

func (txp *TxPool) Flush() {
	txp.transactions = make(map[[32]uint8]*core.Transaction)
}

func (txp *TxPool) TxMapSorterFunc() *TxMapSorter {
	txx := make([]*core.Transaction, len(txp.transactions))

	i := 0
	for _, val := range txp.transactions {
		txx[i] = val
		i++
	}

	s := &TxMapSorter{txx}

	sort.Sort(s)

	return s
}

func (s *TxMapSorter) Len() int { return len(s.transactions) }

func (s *TxMapSorter) Swap(i, j int) {
	s.transactions[i], s.transactions[j] = s.transactions[j], s.transactions[i]
}

func (s *TxMapSorter) Less(i, j int) bool {
	return s.transactions[i].FirstSeen().Before(s.transactions[j].FirstSeen())
}

package network

import (
	"strconv"
	"testing"
	"time"

	"github.com/bansaltushar014/go-blockchain-l2/core"
	"github.com/stretchr/testify/assert"
)

func TestNewTxPool(t *testing.T) {
	txp := NewTxPool()
	assert.NotEmpty(t, txp)
}

func TestAddTransactionsToPool(t *testing.T) {
	txp := NewTxPool()
	assert.NotEmpty(t, txp)
	t1 := core.Trans{}
	txx := t1.CreateRandomTx([]byte("First Transaction"))
	txx.SetFirstSeen(time.Now())
	txp.Add(txx)
	assert.Equal(t, txp.Len(), 1)
	hash := txx.Hash(txx.Data)
	ok := txp.Has(hash)
	assert.True(t, ok)
	txp.Flush()
	assert.Equal(t, txp.Len(), 0)
}

func TestSortTransactions(t *testing.T) {
	p := NewTxPool()
	txLen := 10
	t1 := core.Trans{}
	for i := 0; i < txLen; i++ {
		str := "First Transaction " + strconv.Itoa(i)
		txx := t1.CreateRandomTx([]byte(str))
		txx.SetFirstSeen(time.Now())
		p.Add(txx)
	}
	assert.Equal(t, 10, p.Len())
	txx := p.TxMapSorterFunc()
	for i := 0; i < txLen-1; i++ {
		assert.True(t, txx.transactions[i].FirstSeen().Before(txx.transactions[i+1].FirstSeen()))
	}
}

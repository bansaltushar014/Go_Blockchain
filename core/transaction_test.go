package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSignAndVerifyTxn(t *testing.T) {
	tx := NewTransaction()
	tx.SetFirstSeen(time.Now())
	assert.NotEmpty(t, tx)
	bool := tx.Verify()
	assert.True(t, bool)
}

func TestCreateTxAndVerify(t *testing.T) {
	t1 := Trans{}
	txx := t1.CreateRandomTx([]byte("First Transaction"))
	txx.SetFirstSeen(time.Now())
	assert.NotEmpty(t, txx)
	bool := txx.Verify()
	assert.True(t, bool)

}
